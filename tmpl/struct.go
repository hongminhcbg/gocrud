package tmpl

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var userValidate *validator.Validate

func init() {
	userValidate = validator.New()
}

// template for struct

type User struct {
	Id        int64      `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Age       string     `json:"age,omitempty"`
	Email     string     `json:"email,omitempty"`
	Sex       bool       `json:"sex,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (User) TableName() string {
	return "users"
}

func (m *User) Validate() error {
	return userValidate.Struct(m)
}
