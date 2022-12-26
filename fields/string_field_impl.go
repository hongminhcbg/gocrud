package fields

import (
	"fmt"

	"github.com/hongminhcbg/gocrud/utils"
)

type _string struct {
	nameSnakeCase string
	nameCamelCase string
	comment       string
}

func NewFieldString(name, comment string) IField {
	return &_string{
		nameSnakeCase: name,
		nameCamelCase: utils.SnakeToCamel(name),
		comment:       comment,
	}
}

func (s *_string) Name() string {
	return s.nameCamelCase
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
