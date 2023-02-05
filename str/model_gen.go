package str

import (
	"fmt"
	"os"
	"strings"
)

var headerFile = `
package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-gorm/datatypes"
)

var %sValidate *validator.Validate

func init() {
	%sValidate = validator.New()
}

func (m *%s) Validate() error {
	return %sValidate.Struct(m)
}

type %s struct {
`

func (s *Struct) GenModelFile() error {
	ans := new(strings.Builder)
	_, err := ans.WriteString(fmt.Sprintf(headerFile, s.Name, s.Name, s.NameCamel, s.Name, s.NameCamel))
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
	_, err = ans.WriteString(fmt.Sprintf("func (%s) TableName() string {\n  return \"%s\"\n}", s.NameCamel, strings.ToLower(s.NameCamel)))
	fileName := fmt.Sprintf("%s_model.go", s.Name)
	os.Remove(fileName)
	return os.WriteFile(fileName, []byte(ans.String()), 0644)
}
