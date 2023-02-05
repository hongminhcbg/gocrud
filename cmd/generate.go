package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/hongminhcbg/gocrud/fields"
	"github.com/hongminhcbg/gocrud/str"
	"github.com/hongminhcbg/gocrud/utils"
	"github.com/urfave/cli/v3"
	"gopkg.in/yaml.v2"
)

type Field struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Validate string `yaml:"validate,omitempty"`
	Comment  string `yaml:"comment"`
	SqlHint  string `yaml:"sql_hint"`
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
		sqlHint := utils.ParseSqlHint(c.Fields[i].SqlHint)
		fmt.Println("parse sql hint", " fieldName:", c.Fields[i].Name, " hint:", c.Fields[i].SqlHint, " output is ", sqlHint)
		if strings.EqualFold(c.Fields[i].Type, "text") {
			fieldList = append(fieldList, fields.NewFieldString(c.Fields[i].Name, c.Fields[i].Comment, sqlHint, c.Fields[i].Validate))
			continue
		}

		if strings.EqualFold(c.Fields[i].Type, "int8") {
			fieldList = append(fieldList, fields.NewInt(c.Fields[i].Name, c.Fields[i].Comment, fields.Int8, sqlHint, c.Fields[i].Validate))
			continue
		}

		if strings.EqualFold(c.Fields[i].Type, "int16") {
			fieldList = append(fieldList, fields.NewInt(c.Fields[i].Name, c.Fields[i].Comment, fields.Int16, sqlHint, c.Fields[i].Validate))
			continue
		}

		if strings.EqualFold(c.Fields[i].Type, "int32") {
			fieldList = append(fieldList, fields.NewInt(c.Fields[i].Name, c.Fields[i].Comment, fields.Int32, sqlHint, c.Fields[i].Validate))
			continue
		}

		if strings.EqualFold(c.Fields[i].Type, "int64") {
			fieldList = append(fieldList, fields.NewInt(c.Fields[i].Name, c.Fields[i].Comment, fields.Int64, sqlHint, c.Fields[i].Validate))
			continue
		}

		if strings.EqualFold(c.Fields[i].Type, "bool") {
			fieldList = append(fieldList, fields.NewBoolField(c.Fields[i].Name, c.Fields[i].Comment, sqlHint, c.Fields[i].Validate))
		}

		if strings.EqualFold(c.Fields[i].Type, "json") {
			fieldList = append(fieldList, fields.NewJsonFileld(c.Fields[i].Name, c.Fields[i].Comment, c.Fields[i].Validate, sqlHint))
		}
	}

	stru := str.NewStruct(c.Name, fieldList)
	wg := new(sync.WaitGroup)
	wg.Add(3)
	go func() {
		err := stru.GenModelFile()
		if err != nil {
			log.Println(err)
		}
		wg.Done()
	}()

	go func() {
		err := stru.GenMigrateFile()
		if err != nil {
			log.Println(err)
		}
		wg.Done()
	}()
	go func() {
		err := stru.GenStoreFile()
		if err != nil {
			log.Println(err)
		}
		wg.Done()
	}()

	wg.Wait()
	return nil
}
