package fields

import (
	"fmt"

	"github.com/hongminhcbg/gocrud/utils"
)

type _bool struct {
	snakeCase string
	camelCase string
	comment   string
}

func NewBoolField(snakeCase, comment string) IField {
	return &_bool{
		snakeCase: snakeCase,
		camelCase: utils.SnakeToCamel(snakeCase),
		comment:   comment,
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
	return fmt.Sprintf("`gorm:\"column:%s\" json:\"%s,omitempty\"`", b.snakeCase, b.snakeCase)
}

func (b *_bool) Comment() string {
	return b.comment
}
