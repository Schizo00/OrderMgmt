package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Customer_ID uint
	Name        string
	Email       string
}
