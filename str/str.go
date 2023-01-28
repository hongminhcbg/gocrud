package str

import (
	"fmt"
	"os"
	"strings"

	"github.com/hongminhcbg/gocrud/fields"
	"github.com/hongminhcbg/gocrud/utils"
)

type Struct struct {
	NameSnake string
	Fields    []fields.IField
	Name      string
}

func NewStruct(name string, fields []fields.IField) *Struct {
	return &Struct{
		NameSnake: utils.SnakeToCamel(name),
		Name:      name,
		Fields:    fields,
	}
}

func (s *Struct) GenModelFile() error {
	ans := new(strings.Builder)
	_, err := ans.WriteString(fmt.Sprintf("package YOUR_PACKAGE_HERE\ntype %s struct {\n", s.NameSnake))
	if err != nil {
		return err
	}
	for i := 0; i < len(s.Fields); i++ {
		_, err := ans.WriteString(fmt.Sprintf("  %s \t%s \t%s // %s\n", s.Fields[i].Name(), s.Fields[i].DataType(), s.Fields[i].Annotation(), s.Fields[i].Comment()))
		if err != nil {
			return err
		}
	}

	ans.WriteString("}\n")
	_, err = ans.WriteString(fmt.Sprintf("func (%s) TableName() string {\n  return \"%s\"\n}", s.NameSnake, strings.ToLower(s.NameSnake)))
	fileName := fmt.Sprintf("%s_model.go", s.Name)
	os.Remove(fileName)
	return os.WriteFile(fileName, []byte(ans.String()), 0644)
}

func (s *Struct) GenMigrateFile() error {
	if len(s.Fields) == 0 {
		return nil
	}

	ans := new(strings.Builder)
	ans.WriteString(fmt.Sprintf("CREATE TABLE `%s`(\n", s.Name))
	for i := 0; i < len(s.Fields)-1; i++ {
		ans.WriteString(fmt.Sprintf("\t%s,\n", s.Fields[i].GenSql()))
	}

	ans.WriteString(fmt.Sprintf("\t%s\n", s.Fields[len(s.Fields)-1].GenSql()))
	ans.WriteString(fmt.Sprintf(") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT '%s'; ", s.Name))
	//log.Println("migrate with ", ans.String())
	fileName := fmt.Sprintf("%s_migrate.sql", s.Name)
	os.Remove(fileName)
	return os.WriteFile(fileName, []byte(ans.String()), 0644)
}

func GenSQL(f fields.IField) string {
	return ""
}
