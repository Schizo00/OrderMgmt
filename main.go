package main

import (
	"OrderMgmt/controllers"
	"OrderMgmt/initializers"

	// "OrderMgmt/models"
	"fmt"

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

	// initializers.DB.Migrator().DropTable(&models.Customer{})
	// initializers.DB.Migrator().DropTable(&models.Order{})
	// initializers.DB.Migrator().DropTable(&models.ProductOrder{})
	// initializers.DB.Migrator().DropTable(&models.Product{})
	fmt.Println("Build Done")
	r := gin.Default()

	r.POST("/products/create", controllers.CreateProduct)
	r.GET("/products/retrieve", controllers.RetrieveAllProducts)
	r.GET("products/retrieve/:index", controllers.RetrieveProductByIndex)
	r.PUT("products/update/:index", controllers.UpdateProductByIndex)
	r.DELETE("products/delete/:index", controllers.DeleteProductByIndex)

	r.POST("/orders/create", controllers.CreateOrder)
	r.GET("/orders/retrieve", controllers.RetrieveAllOrders)
	r.GET("/orders/retrieve/:index", controllers.RetrieveOrderByIndex)
	r.PUT("/orders/update/:index", controllers.UpdateOrderByIndex)
	r.DELETE("/orders/delete/:index", controllers.DeleteOrderByIndex)

	r.POST("/customers/create", controllers.CreateCustomer)
	r.GET("/customers/retrieve", controllers.RetrieveAllCustomers)
	r.GET("/customers/retrieve/:index", controllers.RetrieveCustomerByIndex)
	r.PUT("/customers/update/:index", controllers.UpdateCustomerByIndex)
	r.DELETE("/customers/delete/:index", controllers.DeleteCustomerByIndex)

	r.POST("/productorder/create", controllers.CreateProductOrder)
	r.GET("/productorder/retrieve", controllers.RetrieveAllProductOrders)
	r.GET("/productorder/retrieve/:index", controllers.RetrieveProductOrderByIndex)
	r.PUT("/productorder/update/:index", controllers.UpdateProductOrderByIndex)
	r.DELETE("/productorder/delete/:index", controllers.DeleteProductOrderByIndex)
	r.Run()

}
