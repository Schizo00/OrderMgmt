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

	initializers.DB.Migrator().CreateConstraint(&models.Customer{}, "orders")
	initializers.DB.Migrator().CreateConstraint(&models.Customer{}, "fk_orders_customers")

	initializers.DB.Migrator().CreateConstraint(&models.ProductOrder{}, "customers")
	initializers.DB.Migrator().CreateConstraint(&models.ProductOrder{}, "fk_product_orders_customers")

	initializers.DB.Migrator().CreateConstraint(&models.ProductOrder{}, "orders")
	initializers.DB.Migrator().CreateConstraint(&models.ProductOrder{}, "fk_product_orders_orders")

	initializers.DB.Migrator().CreateConstraint(&models.ProductOrder{}, "products")
	initializers.DB.Migrator().CreateConstraint(&models.ProductOrder{}, "fk_products_customers")

	initializers.DB.Migrator().HasConstraint(&models.ProductOrder{}, "customers")
	initializers.DB.Migrator().HasConstraint(&models.ProductOrder{}, "fk_orders_customers")

}
