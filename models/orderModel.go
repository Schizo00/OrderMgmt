package models

type Order struct {
	OrderID       int `gorm:"primary_key; auto_increment; not null"`
	CustID        int
	ProductOrders []ProductOrder `gorm:"ForeignKey:OrderID"`
	Total         int
}
