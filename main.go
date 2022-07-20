package main

import (
	"OrderMgmt/controllers"
	"OrderMgmt/initializers"

	// "OrderMgmt/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	r.POST("/", controllers.CreateProduct)
	r.GET("/", controllers.RetrieveAllProducts)
	r.GET("/:index", controllers.RetrieveProductByIndex)
	r.Run()

}
