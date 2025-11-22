package models

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	Name    string `json:"name"`
	Email   string `json:"email" gorm:"unique"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
