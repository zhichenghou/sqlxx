# sqlxx

a tool to generate mybatis common code.

go install and run sqlxx

```shell
Usage:
  -db string
    	db conn str: e.g. user:password@tcp(127.0.0.1:3306) (default "")
  -mapperpkg string
    	mapper package name: e.g. com.houzhicheng.mapper
  -modelpkg string
    	model package name: e.g. com.houzhicheng.model
  -schema string
    	selected schema: e.g. test (default "test")
  -table string
    	selected table: e.g. test
```