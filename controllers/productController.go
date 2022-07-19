package controllers

import (
	"OrderMgmt/initializers"
	"OrderMgmt/models"
	"net/http"

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

	c.JSON(http.StatusOK, gin.H{
		"Product": temp_product,
	})
}
