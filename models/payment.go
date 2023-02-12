package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	PatientId     int  `json:"patientId" form:"patientId"`
	DateOfPayment int  `json:"dop" form:"dop"`
	TotalCharge   int  `json:"totalCharge" form:"totalCharge"`
	PaymentStatus bool `json:"paymentStatus" form:"paymentStatus"`

	Patient Patient `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Payment) TableName() string {
	return "payments"
}
