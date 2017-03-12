package models

import (
	"database/sql"
)

type Column struct {
	FieldOrdinal int
	ColumnName   string
	DataType     string
	IsPrimaryKey bool
}

func GetColumns(db *sql.DB, schema string, tableName string) ([]*Column, error) {
	const sqlstr = `SELECT ` +
		`ordinal_position AS field_ordinal, ` +
		`column_name, ` +
		`IF(data_type = 'enum', column_name, column_type) AS data_type, ` +
		`IF(column_key = 'PRI', true, false) AS is_primary_key ` +
		`FROM information_schema.columns ` +
		`WHERE table_schema = ? AND table_name = ? ` +
		`ORDER BY ordinal_position`

	rows, err := db.Query(sqlstr, schema, tableName)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	res := []*Column{}
	for rows.Next() {
		column := Column{}

		err := rows.Scan(&column.FieldOrdinal, &column.ColumnName, &column.DataType, &column.IsPrimaryKey)
		if err != nil {
			return nil, err
		}

		res = append(res, &column)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return res, nil
}
