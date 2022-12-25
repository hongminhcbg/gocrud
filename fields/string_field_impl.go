package fields

import (
	"fmt"
	"strings"
)

type _string struct {
	nameSnakeCase string
	nameCamelCase string
}

func NewFieldString(name string) IField {
	return &_string{
		nameSnakeCase: name,
		nameCamelCase: strings.ToTitle(name),
	}
}

func (s *_string) Name() string {
	return s.nameSnakeCase
}

func (s *_string) DataType() string {
	return "string"
}

func (s *_string) Annotation() string {
	return fmt.Sprintf("`gorm:\"column:%s\" json:\"%s,omitempty\"`", s.nameSnakeCase, s.nameSnakeCase)
}
