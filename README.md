# Public BI benchmark for Tidb

There are 46 workbooks containing 206 tables (.csv files) with the total size of 41 GB compressed and 386 GB uncompressed.

## Quick start
```
go run main.go prepare -h 127.0.0.1 -P 4000 -u root -d test > log.txt
go run main.go run -h 127.0.0.1 -P 4000 -u root -d test > log.txt
```

# Usage
```
go run main.go [prepare|run|cleanup]:
-P int
    	MySQL Port (default 3306)
  -c value
    	test case you want to benchmark
  -d string
    	MySQL database (default "test")
  -h string
    	MySQL Host (default "127.0.0.1")
  -i string
    	input benchmark resource (default "./benchmark")
  -ignore-run-error
    	Ignore errors when run
  -o string
    	output directory to save cvs files and logs (default "./var")
  -p string
    	MySQL Password
  -sample
    	Use sample to test
  -u string
    	MySQL User (default "root")
```

# Now status
To run Bibenchmark in tidb, I try to delate some test cases that is not support now.

Maybe the test cases will be added in future.
## Remove database:
'Wins' and 'USCensus'

## Change test cases:
bigint to signed: Too many files.

## Remove test cases:
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