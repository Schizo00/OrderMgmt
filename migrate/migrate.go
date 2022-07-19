package main

import (
	"OrderMgmt/initializers"
	"OrderMgmt/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Product{})
	initializers.DB.AutoMigrate(&models.Order{})
	initializers.DB.AutoMigrate(&models.ProductOrder{})
	initializers.DB.AutoMigrate(&models.Customer{})
}
