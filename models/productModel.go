package models

type Product struct {
	ProductID     int `gorm:"primary_key; auto_increment; not null"`
	Name          string
	ProductOrders []ProductOrder `gorm:"ForeignKey:ProductID"`
	Price         int
}
