package models

import "gorm.io/gorm"

type Publisher struct {
	gorm.Model
	Name    string `json:"name"`
	Website string `json:"website"`
	Books   []Book `gorm:"foreignKey:PublisherID"`
}
