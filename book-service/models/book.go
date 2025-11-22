package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title           string `json:"title"`
	ISBN            string `json:"isbn" gorm:"unique"`
	PublishedYear   int    `json:"published_year"`
	TotalCopies     int    `json:"total_copies"`
	AvailableCopies int    `json:"available_copies"`

	AuthorID uint   `json:"author_id"`
	Author   Author `gorm:"foreignKey:AuthorID"`

	CategoryID uint     `json:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryID"`

	PublisherID uint      `json:"publisher_id"`
	Publisher   Publisher `gorm:"foreignKey:PublisherID"`
}
