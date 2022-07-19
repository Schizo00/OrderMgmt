package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Order_ID    uint
	Customer_ID uint
	Total       uint
}
