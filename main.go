package main

import (
	"OrderMgmt/controllers"
	"OrderMgmt/initializers"

	// "OrderMgmt/models"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	ID    uint
	Code  string
	Price uint
}

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	fmt.Println("Build Done")

	r := gin.Default()
	r.Use(cors.Default())
	api := r.Group("/api")
	{
		api.GET("/ping", controllers.Ping)
		api.POST("/products/create", controllers.CreateProduct)
		api.GET("/products/retrieve", controllers.RetrieveAllProducts)
		api.GET("/products/retrieve/:index", controllers.RetrieveProductByIndex)
		api.PUT("/products/update/:index", controllers.UpdateProductByIndex)
		api.DELETE("/products/delete/:index", controllers.DeleteProductByIndex)

		api.POST("/orders/create", controllers.CreateOrder)
		api.GET("/orders/retrieve", controllers.RetrieveAllOrders)
		api.GET("/orders/retrieve/:index", controllers.RetrieveOrderByIndex)
		api.PUT("/orders/update/:index", controllers.UpdateOrderByIndex)
		api.DELETE("/orders/delete/:index", controllers.DeleteOrderByIndex)

		api.POST("/customers/create", controllers.CreateCustomer)
		api.GET("/customers/retrieve", controllers.RetrieveAllCustomers)
		api.GET("/customers/retrieve/:index", controllers.RetrieveCustomerByIndex)
		api.PUT("/customers/update/:index", controllers.UpdateCustomerByIndex)
		api.DELETE("/customers/delete/:index", controllers.DeleteCustomerByIndex)

		api.POST("/productorders/create", controllers.CreateProductOrder)
		api.GET("/productorders/retrieve", controllers.RetrieveAllProductOrders)
		api.GET("/productorders/retrieve/:index", controllers.RetrieveProductOrderByIndex)
		api.PUT("/productorders/update/:index", controllers.UpdateProductOrderByIndex)
		api.DELETE("/productorders/delete/:index", controllers.DeleteProductOrderByIndex)

		api.GET("/refresh", controllers.RefreshTotal)

	}

	r.Run()

}
