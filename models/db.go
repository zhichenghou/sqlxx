package models

import (
	"database/sql"
)

type Table struct {
	Schema    string
	TableName string
}

func GetTables(db *sql.DB, schema string) ([]*Table, error) {
	const sqlstr = `SELECT ` +
		`table_schema, ` +
		`table_name ` +
		`FROM information_schema.tables ` +
		`WHERE table_schema = ? AND table_type = 'BASE TABLE'`

	rows, err := db.Query(sqlstr, schema)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tables := []*Table{}
	for rows.Next() {
		table := Table{}
		err := rows.Scan(&table.Schema, &table.TableName)
		if err != nil {
			return nil, err
		}
		tables = append(tables, &table)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return tables, nil
}
