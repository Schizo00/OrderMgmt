package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	ID    uint
	Name  string
	Email string
}
