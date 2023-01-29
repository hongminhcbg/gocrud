package str

import (
	"fmt"
	"os"
	"strings"

	"github.com/hongminhcbg/gocrud/fields"
	"github.com/hongminhcbg/gocrud/utils"
)

type Struct struct {
	NameCamel string
	Fields    []fields.IField
	Name      string
}

func NewStruct(name string, fields []fields.IField) *Struct {
	return &Struct{
		NameCamel: utils.SnakeToCamel(name),
		Name:      name,
		Fields:    fields,
	}
}

func (s *Struct) GenModelFile() error {
	ans := new(strings.Builder)
	_, err := ans.WriteString(fmt.Sprintf("package main\ntype %s struct {\n", s.NameCamel))
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

var formatStore = `
package main

import (
	"context"

	"gorm.io/gorm"
)

type %sStore struct {
	*gorm.DB
}

func new%sStore(db *gorm.DB) *%sStore {
	return &%sStore{db}
}

func (s *%sStore) Create(ctx context.Context, r *%s) error {
	return s.DB.WithContext(ctx).Create(r).Error
}

func (s *%sStore) Save(ctx context.Context, r *%s) error {
	return s.DB.WithContext(ctx).Save(r).Error
}
	`

func (s *Struct) GenStoreFile() error {
	ans := new(strings.Builder)
	header := fmt.Sprintf(formatStore, s.NameCamel, s.NameCamel, s.NameCamel, s.NameCamel, s.NameCamel, s.NameCamel, s.NameCamel, s.NameCamel)
	ans.WriteString(header)

	fileName := fmt.Sprintf("%s_store.go", s.Name)
	os.Remove(fileName)
	return os.WriteFile(fileName, []byte(ans.String()), 0644)
}
