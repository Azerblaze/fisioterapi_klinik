package models

import (
	"gorm.io/gorm"
)

type Treatment struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Cost        int    `json:"cost" form:"cost"`
	CostForeign int    `json:"costForeign" form:"costForeign"`
}

func (Treatment) TableName() string {
	return "treatments"
}
