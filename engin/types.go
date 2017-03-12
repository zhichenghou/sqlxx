package engin

import (
	"log"
	"regexp"
)

func SqlTypeToJavaType(sqlType string) string {
	// ignore precision
	var PrecScaleRE = regexp.MustCompile(`\(([0-9]+)(\s*,[0-9]+)?\)$`)
	m := PrecScaleRE.FindStringSubmatchIndex(sqlType)
	if m != nil {
		sqlType = sqlType[:m[0]] + sqlType[m[1]:]
	}

	var javaType string

	switch sqlType {
	case "bool", "boolean":
		javaType = "Boolean"

	case "char", "varchar", "tinytext", "text", "mediumtext", "longtext":
		javaType = "String"

	case "bit", "tinyint", "smallint", "mediumint", "int", "integer":
		javaType = "Integer"

	case "bigint":
		javaType = "Long"

	case "float":
		javaType = "Float"

	case "decimal", "double":
		javaType = "Double"

	case "binary", "varbinary", "tinyblob", "blob", "mediumblob", "longblob":
		javaType = "String"

	case "timestamp", "datetime", "date", "time":
		javaType = "Date"

	default:
		log.Fatal("unknow type " + sqlType)

	}
	return javaType
}
