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
	return os.WriteFile(fmt.Sprintf("%s_model.go", s.Name), []byte(ans.String()), 0644)
}
