package models

import "gorm.io/gorm"

type ProductOrder struct {
	gorm.Model
	ID uint

	Product   Product `gorm:"foreignKey:ProductID; references:ID"`
	ProductID uint

	Order   Order `gorm:"foreignKey:OrderID; references:ID"`
	OrderID uint

	Customer   Customer `gorm:"foreignKey:CustomerID; references:ID"`
	CustomerID uint

	Quantity uint
	Price    uint
	SubTotal uint
}
