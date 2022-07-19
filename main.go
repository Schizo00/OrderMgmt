package main

import (
	"OrderMgmt/controllers"
	"OrderMgmt/initializers"
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
	fmt.Println("Build Done")
	r := gin.Default()

	r.POST("/", controllers.CreateProduct)
	r.Run()

}
