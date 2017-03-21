package engin

import (
	"database/sql"
	"github.com/zhichenghou/sqlxx/models"
	"os"
	"path"
	"runtime"
	"text/template"
)

type ColumnItem struct {
	SqlName  string
	SqlType  string
	PropName string
	PropType string
}

type Param struct {
	Schema          string
	Table           string
	BasePackageName string
	ModelClassName  string
	ModelParamName  string
	ColumnItems     []*ColumnItem
}

func GenBasic(db *sql.DB, schema string, table string, pkg string) error {
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
		Schema:          schema,
		Table:           table,
		BasePackageName: pkg,
		ModelClassName:  modelClassName,
		ModelParamName:  modelParamName,
		ColumnItems:     columnItems,
	}

	println(modelClassName + "Mapper.xml")
	genMapperXml(&param)

	println(modelClassName + ".java")
	genModelJava(&param)

	println(modelClassName + "Mapper.java")
	genMapperJava(&param)

	println(modelClassName + "Repo.java")
	genRepoJava(&param)

	return nil
}

func genRepoJava(param *Param) error {
	templateDir := getTemplateDir()
	t := template.Must(template.ParseFiles(templateDir + "/java/repo.java.tpl"))
	err := t.ExecuteTemplate(os.Stdout, "repo", param)
	if err != nil {
		return err
	}

	return nil
}

func genModelJava(param *Param) error {
	templateDir := getTemplateDir()
	t := template.Must(template.ParseFiles(templateDir + "/java/model.java.tpl"))
	err := t.ExecuteTemplate(os.Stdout, "model", param)
	if err != nil {
		return err
	}

	return nil
}

func genMapperJava(param *Param) error {
	templateDir := getTemplateDir()
	t := template.Must(template.ParseFiles(templateDir + "/java/mapper.java.tpl"))
	err := t.ExecuteTemplate(os.Stdout, "mapper", param)
	if err != nil {
		return err
	}

	return nil
}

func genMapperXml(param *Param) error {
	templateDir := getTemplateDir()
	t := template.Must(template.ParseFiles(
		templateDir+"/xml/header.xml.tpl",
		templateDir+"/xml/footer.xml.tpl",
		templateDir+"/xml/body.xml.tpl"))
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

func getTemplateDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Dir(filename) + "/../templates"

}
