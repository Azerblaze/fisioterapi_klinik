package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" validate:"required,min=4"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
	Phone    string `json:"phone" form:"phone" validate:"required,alphanum"`
	Address  string `json:"address" form:"address" validate:"required,min=10"`
	Role     int    `json:"role" form:"role" validate:"required,numeric"`
	Status   bool   `json:"status" form:"status"`
	IsAdmin  bool   `json:"isAdmin" form:"isAdmin"`
}

func (User) TableName() string {
	return "users"
}
