package controllers

import (
	"OrderMgmt/initializers"
	"OrderMgmt/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {

	var product struct {
		Name          string `json:"Name"`
		Price         int    `json:"Price"`
		ProductOrders []models.ProductOrder
	}

	c.Bind(&product)

	temp_product := models.Product{Name: product.Name, Price: product.Price, ProductOrders: product.ProductOrders}

	result := initializers.DB.Create(&temp_product)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"Product Name":  product.Name,
		"Product Price": product.Price,
	})
}

func RetrieveAllProducts(c *gin.Context) {
	//var records []models.Product
	var product []models.Product
	initializers.DB.First(&product, 1)

	c.JSON(200, gin.H{
		"Products": product,
	})
}

func RetrieveProductByIndex(c *gin.Context) {
	var product models.Product
	index := c.Param("index")
	initializers.DB.Find(&product, index)
	fmt.Println(index)
	c.JSON(200, gin.H{
		"Product": product,
	})
}

func UpdateProductByIndex(c *gin.Context) {
	var product models.Product
	var userInput struct {
		Name          string `json:"Name"`
		Price         int    `json:"Price"`
		ProductOrders []models.ProductOrder
	}
	c.Bind(&userInput)

	index := c.Param("index")
	initializers.DB.First(&product, index)
	if userInput.Name != "" {
		initializers.DB.Model(&product).Update(models.Product{Name: userInput.Name})
	}

	if userInput.Price != 0 {
		initializers.DB.Model(&product).Update(models.Product{Price: userInput.Price})
	}

	c.JSON(200, gin.H{
		"Product": product,
	})
}

func DeleteProductByIndex(c *gin.Context) {
	index := c.Param("index")
	initializers.DB.Delete(&models.Product{}, index)
	fmt.Println(index)
	c.Status(200)
}
