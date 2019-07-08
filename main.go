package main

import (
	"bufio"
	"compress/bzip2"
	"database/sql"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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

	keepCSV   bool
	newCSV    bool
	useSample bool

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

	cmdLine.BoolVar(&keepCSV, "keep-csv", true, "Keep CVS files when cleanup")
	cmdLine.BoolVar(&newCSV, "new-csv", false, "Download CVS file forcily when prepare")
	cmdLine.BoolVar(&useSample, "sample", false, "Use sample to test")

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

func downloadFile(filePath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	basePath := path.Dir(filePath)
	os.MkdirAll(basePath, 0755)

	if fileExists(filePath) && !newCSV {
		return nil
	}

	os.Remove(filePath)

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	cr := bzip2.NewReader(resp.Body)
	br := bufio.NewReaderSize(cr, 16*1024)
	_, err = io.Copy(out, br)
	return err
}

func download(name string) error {
	fileName := path.Join(input, name, "data-urls.txt")
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	urls := strings.Split(strings.TrimSpace(string(data)), "\n")
	for _, url := range urls {
		seps := strings.Split(url, "/")
		baseName := strings.TrimRight(seps[len(seps)-1], ".bz2")
		fileName = path.Join(output, name, "csv", baseName)

		start := time.Now()
		fmt.Printf("begin to download %s\n", url)
		if err = downloadFile(fileName, url); err != nil {
			return err
		}
		fmt.Printf("download %s successfully, takes %s\n", url, time.Now().Sub(start))
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
		_, err := db.Exec("LOAD DATA LOCAL INFILE '" + filePath + "' INTO TABLE " + tableName)
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

		query := string(data)
		start := time.Now()
		fmt.Printf("begin to execute %s\n", query)
		rows, err := db.Query(query)
		if err != nil {
			return err
		}
		for rows.Next() {

		}
		// Check for errors from iterating over rows.
		if err := rows.Err(); err != nil {
			return err
		}
		fmt.Printf("execute %s, takes %s\n", query, time.Now().Sub(start))
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

	if !keepCSV {
		os.RemoveAll(path.Join(output, name, "csv"))
	}
	return nil
}
