package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID         uint
	Customer   Customer `gorm:"foreignKey:CustomerID; references:ID"`
	CustomerID uint
	Total      uint
}
