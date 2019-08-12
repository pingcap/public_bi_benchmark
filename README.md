# Public BI benchmark for TiDB

## 数据集大小

There are 46 workbooks containing 206 tables (.csv files) with the total size of 41 GB compressed and 386 GB uncompressed.

## 使用方法
```
go run main.go [download|prepare|run|cleanup]: 
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
Hint：

1.使用Sample时，直接prepare即可，无需download。

2.使用数据集时，推荐下载好后使用-o来指定路径。（从源网站下载可能过慢，可以科学上网或者从内部服务器下载）

## 现在的数据集

为了可以在TiDB上运行，对数据集和 SQL 查询语句进行了调整。

### 删除的数据集:
'Wins' and 'USCensus' （列数大于 512，TiDB暂时不支持）

'Wins' 数据集中共13条查询语句

'USCensus' 数据集中共8条查询语句

### 因语法不支持而更改的 SQL 查询语句:
Change `cast xx as bigint` to `cast xx as signed`

### 因语法不支持而删除的 SQL 查询语句（及文件）:
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