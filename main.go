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
	var mapperpkg string
	var modelpkg string

	flag.StringVar(&dbConn, "db", "", "db conn str: e.g. user:password@tcp(127.0.0.1:3306)")
	flag.StringVar(&schema, "schema", "test", "selected schema: e.g. test")
	flag.StringVar(&table, "table", "", "selected table: e.g. test")
	flag.StringVar(&mapperpkg, "mapperpkg", "", "mapper package name: e.g. com.houzhicheng.mapper")
	flag.StringVar(&modelpkg, "modelpkg", "", "model package name: e.g. com.houzhicheng.model")
	flag.Parse()

	if dbConn == "" || schema == "" || table == "" || mapperpkg == "" || modelpkg == "" {
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

	engin.GenBasic(db, schema, table, mapperpkg, modelpkg)

}
