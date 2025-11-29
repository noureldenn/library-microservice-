package models

import "gorm.io/gorm"

type Borrow struct {
	gorm.Model
	MemberID uint   `json:"member_id"`
	BookID   uint   `json:"book_id"`
	Status   string `json:"status"` // borrowed / returned
}
