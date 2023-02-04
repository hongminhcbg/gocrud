package fields

import (
	"fmt"
	"strings"

	"github.com/hongminhcbg/gocrud/models"
	"github.com/hongminhcbg/gocrud/utils"
)

type IntType uint8

const (
	Int8 IntType = iota + 1
	Int16
	Int32
	Int64
)

type _int struct {
	snakeCase string
	camelCase string
	comment   string
	t         IntType
	sqlHint   *models.SqlHint
	validate  string
}

func NewInt(snakeCase string, comment string, t IntType, sqlHint *models.SqlHint, validate string) IField {
	if sqlHint == nil {
		sqlHint = new(models.SqlHint)
	}

	return &_int{
		snakeCase: snakeCase,
		camelCase: utils.SnakeToCamel(snakeCase),
		comment:   comment,
		t:         t,
		sqlHint:   sqlHint,
		validate:  validate,
	}
}

func (i *_int) Name() string {
	return i.camelCase
}

func (i *_int) NameSnake() string {
	return i.snakeCase
}

func (i *_int) DataType() string {
	switch i.t {
	case Int8:
		return "int8"
	case Int16:
		return "int16"
	case Int32:
		return "int32"
	case Int64:
		return "int64"
	default:
		panic(fmt.Sprintf("unknow type %d", i.t))
	}
}

func (i *_int) Annotation() string {
	if i.validate == "" {
		return fmt.Sprintf("`gorm:\"column:%s\" json:\"%s,omitempty\"`", i.snakeCase, i.snakeCase)
	}

	return fmt.Sprintf("`gorm:\"column:%s\" json:\"%s,omitempty\" validate:\"%s\"`", i.snakeCase, i.snakeCase, i.validate)

}

func (i *_int) Comment() string {
	return i.comment
}

func (i *_int) GenSql() string {
	ans := new(strings.Builder)
	ans.WriteString(fmt.Sprintf("`%s` ", i.snakeCase))
	sqlType := i.sqlHint.DataType
	if strings.TrimSpace(sqlType) == "" {
		switch i.t {
		case Int8:
			sqlType = "TINYINT"
		case Int16:
			sqlType = "SMALLINT"
		case Int32:
			sqlType = "INT"
		case Int64:
			sqlType = "BIGINT"
		}
	}

	sqlType = sqlType + " "

	ans.WriteString(sqlType)
	if i.sqlHint.IsNotNull {
		ans.WriteString("NOT NULL ")
	}

	if i.sqlHint.DefaultVal != "" {
		ans.WriteString(fmt.Sprintf("DEFAULT '%s' ", i.sqlHint.DefaultVal))
	}

	ans.WriteString(fmt.Sprintf("COMMENT '%s' ", i.comment))

	switch i.sqlHint.KeyType {
	case models.SqlKeyType_CANDIDATE:
		{
			ans.WriteString(fmt.Sprintf(",\n\tKEY `idx_%s`(`%s`)", i.snakeCase, i.snakeCase))
		}
	case models.SqlKeyType_UNIQUE:
		{
			ans.WriteString(fmt.Sprintf(",\n\tUNIQUE KEY `idx_%s`(`%s`)", i.snakeCase, i.snakeCase))
		}
	case models.SqlKeyType_PRIMARY:
		{
			ans.WriteString(fmt.Sprintf(",\n\tPRIMARY KEY(`%s`)", i.snakeCase))
		}
	}

	return ans.String()
}
