package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Role     int    `json:"role" form:"role"`
	Status   bool   `json:"status" form:"status"`
	IsAdmin  bool   `json:"isAdmin" form:"isAdmin"`
}

func (User) TableName() string {
	return "users"
}
