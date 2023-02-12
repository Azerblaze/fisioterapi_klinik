package models

import (
	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	UserId    int    `json:"userId" form:"userId"`
	PatientId int    `json:"patientId" form:"patientId"`
	Date      int    `json:"date" form:"date" validate:"required"`
	Note      string `json:"notes" form:"notes"`
	Status    bool   `json:"status" form:"status"`

	User    User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Patient Patient `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Appointment) TableName() string {
	return "appointments"
}
