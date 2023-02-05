package fields

import (
	"fmt"
	"strings"

	"github.com/hongminhcbg/gocrud/models"
	"github.com/hongminhcbg/gocrud/utils"
)

type _json struct {
	snakeCase string
	camelCase string
	comment   string
	sqlHint   *models.SqlHint
	validate  string
}

func NewJsonFileld(snake, comment, validate string, sqlHint *models.SqlHint) IField {
	if sqlHint == nil {
		sqlHint = new(models.SqlHint)
	}

	return &_json{
		snakeCase: snake,
		camelCase: utils.SnakeToCamel(snake),
		validate:  validate,
		sqlHint:   sqlHint,
	}
}

func (j *_json) Name() string {
	return j.camelCase
}

func (j *_json) DataType() string {
	return "datatypes.JSON"
}

func (j *_json) Annotation() string {
	if j.validate == "" {
		return fmt.Sprintf("`gorm:\"column:%s\" json:\"%s,omitempty\"`", j.snakeCase, j.snakeCase)
	}

	return fmt.Sprintf("`gorm:\"column:%s\" json:\"%s,omitempty\" validate:\"%s\"`", j.snakeCase, j.snakeCase, j.validate)
}

func (j *_json) Comment() string { return j.comment }

func (j *_json) NameSnake() string { return j.snakeCase }

func (j *_json) GenSql() string {
	ans := new(strings.Builder)
	ans.WriteString(fmt.Sprintf("`%s` ", j.snakeCase))
	if j.sqlHint.DataType != "" {
		ans.WriteString(fmt.Sprintf("%s ", j.sqlHint.DataType))
	} else {
		ans.WriteString("JSON ")
	}

	if j.sqlHint.IsNotNull {
		ans.WriteString("NOT NULL ")
	}

	if j.sqlHint.DefaultVal != "" {
		ans.WriteString(fmt.Sprintf("DEFAILT '%s' ", j.sqlHint.DefaultVal))
	}

	return ans.String()
}
