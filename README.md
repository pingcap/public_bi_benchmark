# Public BI benchmark for TiDB

## Data Size

There are 46 workbooks containing 206 tables (.csv files) with the total size of 41 GB compressed and 386 GB uncompressed.

## Usage
```
go run main.go [download|prepare|run|cleanup]: 
//options used to connect the database server.
  -P int
    	MySQL Port (default 4000)
  -d string
    	MySQL database (default "test")
  -h string
    	MySQL Host (default "127.0.0.1")
  -u string
    	MySQL User (default "root")
  -p string
    	MySQL Password
//options used to control the test behavior.
  -c value
    	test case you want to benchmark
  -i string
    	input benchmark resource (default "./benchmark")
  -o string
    	output directory to save cvs files and logs (default "./var")
  -sample
    	Use sample to test
  -ignore-run-error
    	Ignore errors when running
```
Hint：

1. When using sample to test, you can run `prepare` directly without `download` .

2. When using datasets to test, download datasets in other ways first, use options `-o` to select file directory. (Maybe the download speed is too slow, VPN or downloading from internal server can work well.) 

## Status

For running the benchmark in TiDB,  some datasets and SQL queries are adjusted. The adjustment is as follows.

### Deleted datasets:
Datasets 'Wins' and 'USCensus' （The number of columns is greater than 512.）are deleted.

* In workbook 'Wins' , there are 13 queries. 

* In workbook 'USCensus' , there are 8 queries.

### Adjusted SQL queries:
Change `cast xx as bigint` to `cast xx as signed`

### Deleted SQL queries:
N*interval:
```
benchmark/CommonGovernment/queries/2.sql
benchmark/CommonGovernment/queries/22.sql
benchmark/CommonGovernment/queries/23.sql
```

trim,splitpart:
```
benchmark/IGlocations1/queries/16.sql
benchmark/Romance/queries/8.sql (only trim)
benchmark/Romance/queries/9.sql (only trim)
```

median:
```
benchmark/PanCreactomy1/queries/2.sql
benchmark/PanCreactomy2/queries/4.sql
```
cast as timestamp:
```
benchmark/RealEstate2/queries/31.sql
```
cast as text:
```
benchmark/TrainsUK1/queries/6.sql
benchmark/TrainsUK1/queries/7.sql
benchmark/Uberlandia/queries/12.sql
```