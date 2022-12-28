package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/hongminhcbg/gocrud/fields"
	"github.com/hongminhcbg/gocrud/str"
	"github.com/urfave/cli/v3"
	"gopkg.in/yaml.v2"
)

type Field struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Validate string `yaml:"validate,omitempty"`
	Comment  string `yaml:"comment"`
}

type Collection struct {
	Name   string   `yaml:"name"`
	Type   string   `yaml:"type"`
	Fields []*Field `yaml:"fields"`
}

func generate(ctx *cli.Context) error {
	fmt.Println("this generate func", fileInput)
	b, err := os.ReadFile(fileInput)
	if err != nil {
		return err
	}

	c := new(Collection)

	err = yaml.Unmarshal(b, c)
	if err != nil {
		return err
	}

	fieldList := make([]fields.IField, 0, len(c.Fields))
	for i := 0; i < len(c.Fields); i++ {
		if strings.EqualFold(c.Fields[i].Type, "text") {
			fieldList = append(fieldList, fields.NewFieldString(c.Fields[i].Name, c.Fields[i].Comment))
			continue
		}

		if strings.EqualFold(c.Fields[i].Type, "int8") {
			fieldList = append(fieldList, fields.NewInt(c.Fields[i].Name, c.Fields[i].Comment, fields.Int8))
			continue
		}

		if strings.EqualFold(c.Fields[i].Type, "int16") {
			fieldList = append(fieldList, fields.NewInt(c.Fields[i].Name, c.Fields[i].Comment, fields.Int16))
			continue
		}

		if strings.EqualFold(c.Fields[i].Type, "int32") {
			fieldList = append(fieldList, fields.NewInt(c.Fields[i].Name, c.Fields[i].Comment, fields.Int32))
			continue
		}

		if strings.EqualFold(c.Fields[i].Type, "int64") {
			fieldList = append(fieldList, fields.NewInt(c.Fields[i].Name, c.Fields[i].Comment, fields.Int64))
			continue
		}

		if strings.EqualFold(c.Fields[i].Type, "bool") {
			fieldList = append(fieldList, fields.NewBoolField(c.Fields[i].Name, c.Fields[i].Comment))
		}

	}

	stru := str.NewStruct(c.Name, fieldList)
	return stru.GenModelFile()
}
