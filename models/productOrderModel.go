package models

type ProductOrder struct {
	ProductOrderID int `gorm:"primary_key; auto_increment; not null"`
	OrderID        int
	Order          Order
	Product        Product
	ProductID      int
	Quantity       int
	Price          int
	SubTotal       int
}
