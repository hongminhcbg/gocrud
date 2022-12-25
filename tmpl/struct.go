package tmpl

import "time"

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
