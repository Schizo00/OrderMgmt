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
	initializers.DB.DropTableIfExists(&models.Customer{}, &models.Product{}, &models.Order{}, &models.ProductOrder{})
	initializers.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Customer{}, &models.Product{}, &models.Order{}, &models.ProductOrder{})

	initializers.DB.Model(&models.Order{}).AddForeignKey("cust_id", "customers(cust_id)", "CASCADE", "CASCADE")
	initializers.DB.Model(&models.ProductOrder{}).AddForeignKey("order_id", "orders(order_id)", "CASCADE", "CASCADE")
	initializers.DB.Model(&models.ProductOrder{}).AddForeignKey("product_id", "products(product_id)", "CASCADE", "CASCADE")
	// initializers.DB.Migrator().CreateConstraint(&models.Customer{}, "orders")
	// initializers.DB.Migrator().CreateConstraint(&models.Customer{}, "fk_orders_customers")

	// initializers.DB.Migrator().CreateConstraint(&models.ProductOrder{}, "customers")
	// initializers.DB.Migrator().CreateConstraint(&models.ProductOrder{}, "fk_product_orders_customers")

	// initializers.DB.Migrator().CreateConstraint(&models.ProductOrder{}, "orders")
	// initializers.DB.Migrator().CreateConstraint(&models.ProductOrder{}, "fk_product_orders_orders")

	// initializers.DB.Migrator().CreateConstraint(&models.ProductOrder{}, "products")
	// initializers.DB.Migrator().CreateConstraint(&models.ProductOrder{}, "fk_products_customers")

	// initializers.DB.Migrator().HasConstraint(&models.ProductOrder{}, "customers")
	// initializers.DB.Migrator().HasConstraint(&models.ProductOrder{}, "fk_orders_customers")

}
