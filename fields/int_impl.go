package fields

import (
	"fmt"

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
}

func NewInt(snakeCase string, comment string, t IntType) IField {
	return &_int{
		snakeCase: snakeCase,
		camelCase: utils.SnakeToCamel(snakeCase),
		comment:   comment,
		t:         t,
	}
}

func (i *_int) Name() string {
	return i.camelCase
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
	return fmt.Sprintf("`gorm:\"column:%s\" json:\"%s,omitempty\"`", i.snakeCase, i.snakeCase)
}

func (i *_int) Comment() string {
	return i.comment
}
