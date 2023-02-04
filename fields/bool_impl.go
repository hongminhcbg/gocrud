package fields

import (
	"fmt"
	"strings"

	"github.com/hongminhcbg/gocrud/models"
	"github.com/hongminhcbg/gocrud/utils"
)

type _bool struct {
	snakeCase string
	camelCase string
	comment   string
	sqlHint   *models.SqlHint
	validate  string
}

func NewBoolField(snakeCase, comment string, sqlHint *models.SqlHint, validate string) IField {
	if sqlHint == nil {
		sqlHint = new(models.SqlHint)
	}

	return &_bool{
		snakeCase: snakeCase,
		camelCase: utils.SnakeToCamel(snakeCase),
		comment:   comment,
		sqlHint:   sqlHint,
		validate:  validate,
	}
}

func (b *_bool) Name() string {
	return b.camelCase
}

func (b *_bool) NameSnake() string {
	return b.snakeCase
}

func (b *_bool) DataType() string {
	return "bool"
}

func (b *_bool) Annotation() string {
	if b.validate == "" {
		return fmt.Sprintf("`gorm:\"column:%s\" json:\"%s,omitempty\"`", b.snakeCase, b.snakeCase)
	}

	return fmt.Sprintf("`gorm:\"column:%s\" json:\"%s,omitempty\" validate:\"%s\"`", b.snakeCase, b.snakeCase, b.validate)
}

func (b *_bool) Comment() string {
	return b.comment
}

func (b *_bool) GenSql() string {
	ans := new(strings.Builder)
	ans.WriteString(fmt.Sprintf("`%s` ", b.snakeCase))
	sqlType := b.sqlHint.DataType
	if strings.TrimSpace(sqlType) == "" {
		sqlType = "TINYINT(1)"
	}

	sqlType = sqlType + " "
	ans.WriteString(sqlType)
	if b.sqlHint.IsNotNull {
		ans.WriteString("NOT NULL ")
	}

	if b.sqlHint.DefaultVal != "" {
		ans.WriteString(fmt.Sprintf("DEFAULT '%s' ", b.sqlHint.DefaultVal))
	}

	ans.WriteString(fmt.Sprintf("COMMENT '%s' ", b.comment))
	return ans.String()
}
