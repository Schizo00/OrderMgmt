package controllers

import (
	"OrderMgmt/initializers"
	"OrderMgmt/models"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {

	var product struct {
		Name  string
		Price uint
	}

	c.Bind(&product)

	temp_product := models.Product{Name: product.Name, Price: product.Price}

	result := initializers.DB.Create(&temp_product)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"Product Name": product.Name,
	})
}

func ReadAllProducts(c *gin.Context) {
	//var records []models.Product
	var product models.Product
	initializers.DB.First(&product, 1)

	c.JSON(200, gin.H{
		"Product": product,
	})
}
