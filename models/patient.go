package models

import (
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	Name          string `json:"name" form:"name" validate:"required,min=4"`
	Email         string `json:"email" form:"email" validate:"required,email"`
	Phone         string `json:"phone" form:"phone" validate:"required,alphanum"`
	Address       string `json:"address" form:"address" validate:"min=10"`
	DateOfBirth   int    `json:"dob" form:"dob" validate:"required"`
	ForeignStatus bool   `json:"foreignStatus" form:"foreignStatus" validate:"required"`
}

func (Patient) TableName() string {
	return "patients"
}
