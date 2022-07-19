package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Product_ID uint
	Name       string
	Price      uint
}
