package models

import (
	"gorm.io/gorm"
)

type Break struct {
	gorm.Model
	UserId    int    `json:"userId" form:"userId"`
	DateFrom  int    `json:"dateFrom" form:"dateFrom" validate:"required"`
	DateUntil int    `json:"dateUntil" form:"dateUntil" validate:"required"`
	Detail    string `json:"detail" form:"detail" validate:"required"`
	Address   string `json:"address" form:"address" validate:"min=10"`
	Status    bool   `json:"status" form:"status"`

	User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Break) TableName() string {
	return "breaks"
}
