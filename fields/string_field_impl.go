package fields

import (
	"fmt"
	"strings"

	"github.com/hongminhcbg/gocrud/models"
	"github.com/hongminhcbg/gocrud/utils"
)

type _string struct {
	nameSnakeCase string
	nameCamelCase string
	comment       string
	sqlHint       *models.SqlHint
}

func NewFieldString(name, comment string, sqlHint *models.SqlHint) IField {
	if sqlHint == nil {
		sqlHint = new(models.SqlHint)
	}

	return &_string{
		nameSnakeCase: name,
		nameCamelCase: utils.SnakeToCamel(name),
		comment:       comment,
		sqlHint:       sqlHint,
	}
}

func (s *_string) Name() string {
	return s.nameCamelCase
}

func (s *_string) NameSnake() string {
	return s.nameSnakeCase
}

func (s *_string) DataType() string {
	return "string"
}

func (s *_string) Annotation() string {
	return fmt.Sprintf("`gorm:\"column:%s\" json:\"%s,omitempty\"`", s.nameSnakeCase, s.nameSnakeCase)
}

func (s *_string) Comment() string {
	return s.comment
}

func (s *_string) GenSql() string {
	ans := new(strings.Builder)
	ans.WriteString(fmt.Sprintf("`%s` ", s.nameSnakeCase))
	sqlType := s.sqlHint.DataType
	if strings.TrimSpace(sqlType) == "" {
		sqlType = "LONGTEXT"
	}

	sqlType = sqlType + " "
	ans.WriteString(sqlType)
	if s.sqlHint.IsNotNull {
		ans.WriteString("NOT NULL ")
	}

	if s.sqlHint.DefaultVal != "" {
		ans.WriteString(fmt.Sprintf("DEFAULT '%s' ", s.sqlHint.DefaultVal))
	}

	ans.WriteString(fmt.Sprintf("COMMENT '%s' ", s.comment))

	switch s.sqlHint.KeyType {
	case models.SqlKeyType_CANDIDATE:
		{
			ans.WriteString(fmt.Sprintf(",\n\tKEY `idx_%s`(`%s`)", s.nameSnakeCase, s.nameSnakeCase))
		}
	case models.SqlKeyType_UNIQUE:
		{
			ans.WriteString(fmt.Sprintf(",\n\tUNIQUE KEY `idx_%s`(`%s`)", s.nameSnakeCase, s.nameSnakeCase))
		}
	case models.SqlKeyType_PRIMARY:
		{
			ans.WriteString(fmt.Sprintf(",\n\tPRIMARY KEY(`%s`)", s.nameSnakeCase))
		}
	}

	return ans.String()
}
