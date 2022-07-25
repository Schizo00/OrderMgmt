package models

type Customer struct {
	CustID int `gorm:"primary_key; auto_increment; not null; unique"`
	Name   string
	Email  string
	Orders []Order `gorm:"ForeignKey:CustID"`
}
