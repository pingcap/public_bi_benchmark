package main

import (
	"compress/bzip2"
	"database/sql"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"runtime/debug"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
)

type setFlags map[string]struct{}

func (i *setFlags) String() string {
	names := make([]string, 0, len(*i))
	for name := range *i {
		names = append(names, name)
	}
	return strings.Join(names, ",")
}

func (i *setFlags) Set(value string) error {
	(*i)[strings.ToLower(value)] = struct{}{}
	return nil
}

var (
	output string
	input  string

	cases setFlags

	host     string
	port     int
	user     string
	password string
	database string

	useSample    bool
	ignoreRunErr bool

	cmdLine *flag.FlagSet

	db *sql.DB
)

func init() {
	cmdLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	cmdLine.StringVar(&output, "o", "./var", "output directory to save cvs files and logs")
	cmdLine.StringVar(&input, "i", "./benchmark", "input benchmark resource")
	cases = make(setFlags)
	cmdLine.Var(&cases, "c", "test case you want to benchmark")

	cmdLine.StringVar(&host, "h", "127.0.0.1", "MySQL Host")
	cmdLine.IntVar(&port, "P", 3306, "MySQL Port")
	cmdLine.StringVar(&user, "u", "root", "MySQL User")
	cmdLine.StringVar(&password, "p", "", "MySQL Password")
	cmdLine.StringVar(&database, "d", "test", "MySQL database")

	cmdLine.BoolVar(&useSample, "sample", false, "Use sample to test")
	cmdLine.BoolVar(&ignoreRunErr, "ignore-run-error", false, "Ignore errors when run")

	cmdLine.Usage = func() {
		fmt.Fprintf(cmdLine.Output(), "Usage of %s [prepare|run|cleanup]:\n", os.Args[0])
		cmdLine.PrintDefaults()
	}
}

func perr(err error) {
	if err == nil {
		return
	}

	fmt.Printf("meet err: %v\n", err)
	debug.PrintStack()
	os.Exit(1)
}

func main() {
	op := os.Args[1]
	switch op {
	case "prepare":
	case "run":
	case "cleanup":
	default:
		cmdLine.Usage()
		os.Exit(1)
	}
	cmdLine.Parse(os.Args[2:])

	var err error
	if len(password) > 0 {
		password = ":" + password
	}
	uri := fmt.Sprintf("%s%s@tcp(%s:%d)/%s?allowAllFiles=true&sql_mode=ANSI_QUOTES", user, password, host, port, database)
	db, err = sql.Open("mysql", uri)
	defer db.Close()

	names := []string{}

	files, err := ioutil.ReadDir(input)
	perr(err)
	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		_, ok := cases[strings.ToLower(f.Name())]
		if len(cases) == 0 || ok {
			names = append(names, f.Name())
		}
	}
	perr(err)

	// TODO: support parallelism
	for _, name := range names {
		switch op {
		case "prepare":
			err = prepare(name)
		case "run":
			err = run(name)
		case "cleanup":
			err = cleanup(name)
		default:
		}

		perr(err)
	}
}

func fileExists(filepath string) bool {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return false
	}

	return true
}

func decompressCSV(csvPath string, dataPath string) error {
	csvFile, err := os.Create(csvPath)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	dataFile, err := os.Open(dataPath)
	if err != nil {
		return err
	}
	defer dataFile.Close()

	cr := bzip2.NewReader(dataFile)

	_, err = io.Copy(csvFile, cr)
	return err
}

func download(name string) error {
	dataUrls := path.Join(input, name, "data-urls.txt")

	dataDir := path.Join(output, name, "data")
	os.MkdirAll(dataDir, 0755)

	cmd := exec.Command("wget", "-c", "-P", dataDir, "-i", dataUrls)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	files, err := ioutil.ReadDir(dataDir)
	if err != nil {
		return err
	}

	csvDir := path.Join(output, name, "csv")
	os.MkdirAll(csvDir, 0755)
	for _, f := range files {
		dataPath := path.Join(dataDir, f.Name())
		csvPath := path.Join(csvDir, strings.TrimRight(f.Name(), ".bz2"))

		decompressCSV(csvPath, dataPath)
	}

	return nil
}

func listTables(name string) ([]string, []string, error) {
	tablesDir := path.Join(input, name, "tables")
	files, err := ioutil.ReadDir(tablesDir)
	if err != nil {
		return nil, nil, err
	}

	tables := []string{}
	tableSQLs := []string{}

	for _, f := range files {
		data, err := ioutil.ReadFile(path.Join(tablesDir, f.Name()))
		if err != nil {
			return nil, nil, err
		}

		sqlData := string(data)
		tableName := strings.TrimRight(f.Name(), ".table.sql")

		tables = append(tables, tableName)
		tableSQLs = append(tableSQLs, sqlData)
	}

	return tables, tableSQLs, nil
}

func loadCSV(name string) error {

	_, err := db.Exec("set global local_infile = 'ON'")
	if err != nil {
		return err
	}

	root := path.Join(output, name, "csv")
	if useSample {
		root = path.Join(input, name, "samples")
	}

	files, err := ioutil.ReadDir(root)
	if err != nil {
		return err
	}

	for _, f := range files {
		filePath := path.Join(root, f.Name())
		mysql.RegisterLocalFile(filePath)

		tableName := strings.TrimRight(f.Name(), ".csv")

		if useSample {
			tableName = strings.TrimRight(tableName, ".sample")
		}

		start := time.Now()
		fmt.Printf("begin to load csv %s into %s\n", filePath, tableName)
		_, err := db.Exec("LOAD DATA LOCAL INFILE '" + filePath + "' INTO TABLE " + tableName +
			" FIELDS TERMINATED BY '|' OPTIONALLY ENCLOSED BY '\"' LINES TERMINATED BY '\\n'")
		mysql.DeregisterLocalFile(filePath)

		if err != nil {
			return err
		}
		fmt.Printf("load csv %s into %s, takes %s\n", filePath, tableName, time.Now().Sub(start))
	}
	return nil
}

func prepare(name string) error {
	// Try cleanup at first
	cleanup(name)

	// Download csv files
	if !useSample {
		if err := download(name); err != nil {
			return err
		}
	}

	tables, tableSQLs, err := listTables(name)
	if err != nil {
		return err
	}
	for i, tableName := range tables {
		fmt.Printf("begin to create table %s\n", tableName)
		if _, err := db.Exec(tableSQLs[i]); err != nil {
			return err
		}
	}

	if err := loadCSV(name); err != nil {
		return err
	}

	return nil
}

func run(name string) error {
	root := path.Join(input, name, "queries")
	files, err := ioutil.ReadDir(root)
	if err != nil {
		return err
	}

	for _, f := range files {
		data, err := ioutil.ReadFile(path.Join(root, f.Name()))
		if err != nil {
			return err
		}

		queryName := path.Join(name, "queries", f.Name())

		query := string(data)
		start := time.Now()
		fmt.Printf("begin to execute query %s\n", queryName)
		rows, err := db.Query(query)
		fmt.Printf("execute %s, takes %s, err %v\n", queryName, time.Now().Sub(start), err)

		if err != nil {
			if ignoreRunErr {
				continue
			}
			return err
		}
		for rows.Next() {

		}
		// Check for errors from iterating over rows.
		if err := rows.Err(); err != nil {
			return err
		}

	}
	return nil
}

func cleanup(name string) error {
	tables, _, err := listTables(name)
	if err != nil {
		return nil
	}

	for _, tableName := range tables {
		fmt.Printf("begin to drop table %s\n", tableName)
		if _, err = db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName)); err != nil {
			return err
		}
	}

	return nil
}
