package engin

import (
	"database/sql"
	"github.com/zhichenghou/sqlxx/models"
	"os"
	"text/template"
)

type ColumnItem struct {
	SqlName  string
	SqlType  string
	PropName string
	PropType string
}

type Param struct {
	Schema            string
	Table             string
	MapperPackageName string
	MapperClassName   string
	ModelPackageName  string
	ModelClassName    string
	ModelParamName    string
	ColumnItems       []*ColumnItem
}

func GenBasic(db *sql.DB, schema string, table string, mapperPackage string, modelPackage string) error {
	columns, err := models.GetColumns(db, schema, table)
	if err != nil {
		return err
	}

	columnItems := []*ColumnItem{}
	for _, column := range columns {
		columnItem := ColumnItem{
			SqlName:  column.ColumnName,
			SqlType:  column.DataType,
			PropName: UnderscoreToCamelcase(column.ColumnName, false),
			PropType: SqlTypeToJavaType(column.DataType),
		}
		columnItems = append(columnItems, &columnItem)
	}

	modelClassName := UnderscoreToCamelcase(table, true)
	sl := CamelcaseToSlice(modelClassName, true)
	modelParamName := sl[len(sl)-1]

	param := Param{
		Schema:            schema,
		Table:             table,
		MapperPackageName: mapperPackage,
		MapperClassName:   modelClassName + "Mapper",
		ModelPackageName:  modelPackage,
		ModelClassName:    modelClassName,
		ModelParamName:    modelParamName,
		ColumnItems:       columnItems,
	}

	println(modelClassName + "Mapper.xml")
	genMapperXml(&param)

	println(modelClassName + ".java")
	genModelJava(&param)

	println(modelClassName + "Mapper.java")
	genMapperJava(&param)

	return nil
}

func genModelJava(param *Param) error {
	t := template.Must(template.ParseFiles("templates/java/model.java.tpl"))
	err := t.ExecuteTemplate(os.Stdout, "model", param)
	if err != nil {
		return err
	}

	return nil
}

func genMapperJava(param *Param) error {
	t := template.Must(template.ParseFiles("templates/java/mapper.java.tpl"))
	err := t.ExecuteTemplate(os.Stdout, "mapper", param)
	if err != nil {
		return err
	}

	return nil
}

func genMapperXml(param *Param) error {
	t := template.Must(template.ParseFiles("templates/xml/header.xml.tpl", "templates/xml/footer.xml.tpl", "templates/xml/body.xml.tpl"))
	err := t.ExecuteTemplate(os.Stdout, "header", param)
	if err != nil {
		return err
	}

	err = t.ExecuteTemplate(os.Stdout, "body", param)
	if err != nil {
		return err
	}

	err = t.ExecuteTemplate(os.Stdout, "footer", param)
	if err != nil {
		return err
	}

	return nil
}
