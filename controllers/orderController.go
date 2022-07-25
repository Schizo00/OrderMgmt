package controllers

import (
	"OrderMgmt/initializers"
	"OrderMgmt/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {

	var order struct {
		CustId        int
		ProductOrders []models.ProductOrder
		Total         int
	}

	c.Bind(&order)

	temp_order := models.Order{CustID: order.CustId, ProductOrders: order.ProductOrders, Total: order.Total}

	result := initializers.DB.Create(&temp_order)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"Customer ID": order.CustId,
		"Total":       order.Total,
	})
}

func RetrieveAllOrders(c *gin.Context) {
	//var records []models.Product
	var orders []models.Order
	initializers.DB.Find(&orders)

	c.JSON(200, gin.H{
		"Orders": orders,
	})
}

func RetrieveOrderByIndex(c *gin.Context) {
	var order models.Order
	index := c.Param("index")
	initializers.DB.Find(&order, index)
	fmt.Println(index)
	c.JSON(200, gin.H{
		"Order": order,
	})
}

func UpdateOrderByIndex(c *gin.Context) {
	var order models.Order
	var userInput models.Order
	c.Bind(&userInput)

	index := c.Param("index")
	initializers.DB.First(&order, index)
	if userInput.CustID != 0 {
		initializers.DB.Model(&order).Update("CustD", userInput.CustID)
	}

	if userInput.Total != 0 {
		initializers.DB.Model(&order).Update("Price", userInput.Total)
	}

	c.JSON(200, gin.H{
		"Order": order,
	})
}

func DeleteOrderByIndex(c *gin.Context) {
	index := c.Param("index")
	initializers.DB.Delete(&models.Order{}, index)
	fmt.Println(index)
	c.Status(200)
}
