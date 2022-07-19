package models

import "gorm.io/gorm"

type ProductOrder struct {
	gorm.Model
	ID          uint
	Product_ID  uint
	Order_ID    uint
	Customer_ID uint
	Quantity    uint
	Price       uint
	SubTotal    uint
}
