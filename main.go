package main

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zhichenghou/sqlxx/engin"
	"github.com/zhichenghou/sqlxx/models"
	"log"
)

func main() {
	var dbConn string
	var schema string
	var table string
	var pkg string

	flag.StringVar(&dbConn, "db", "", "db conn str: e.g. user:password@tcp(127.0.0.1:3306)")
	flag.StringVar(&schema, "schema", "test", "selected schema: e.g. test")
	flag.StringVar(&table, "table", "", "selected table: e.g. test")
	flag.StringVar(&pkg, "pkg", "", `base package name: e.g. com.houzhicheng, and then
	    	model package: com.houzhicheng.domain.model, 
	    	mapper package: com.houzhicheng.infra.persistence.sql.mapper,
	    	repo package: com.houzhicheng.domain.repo`)
	flag.Parse()

	if dbConn == "" || schema == "" || table == "" || pkg == "" {
		log.Println("Usage:")
		flag.PrintDefaults()
		return
	} else {
		log.Println("connect str: " + dbConn + " / " + schema)
	}

	db, err := sql.Open("mysql", dbConn+"/"+schema)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	tables, err := models.GetTables(db, schema)
	if err != nil {
		log.Fatal(err)
	}

	tableNames := map[string]bool{}
	for _, tableName := range tables {
		tableNames[tableName.TableName] = true
	}

	if _, ok := tableNames[table]; !ok {
		log.Println("table not found " + table)
		return
	}

	engin.GenBasic(db, schema, table, pkg)

}
