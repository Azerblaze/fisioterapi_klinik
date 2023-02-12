package models

import (
	"gorm.io/gorm"
)

type MedicalRecord struct {
	gorm.Model
	PatientId int    `json:"patientId" form:"patientId"`
	Diagnosis string `json:"diagnosis" form:"diagnosis"`
	Progress  string `json:"progress" form:"progress"`
	Note      string `json:"notes" form:"notes"`

	Patient Patient `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (MedicalRecord) TableName() string {
	return "medicalRecords"
}
