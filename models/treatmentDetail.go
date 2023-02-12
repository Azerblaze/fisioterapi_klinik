package models

import (
	"gorm.io/gorm"
)

type TreatmentDetail struct {
	gorm.Model
	MedicalRecordId int `json:"medicalRecordId" form:"medicalRecordId"`
	TreatmentId     int `json:"treatmentId" form:"treatmentId"`

	MedicalRecord MedicalRecord `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Treatment     Treatment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (TreatmentDetail) TableName() string {
	return "treatmentDetails"
}
