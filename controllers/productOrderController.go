package controllers

import (
	"OrderMgmt/initializers"
	"OrderMgmt/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateProductOrder(c *gin.Context) {

	var productOrder struct {
		OrderID   int
		Quantity  int
		Order     models.Order
		Product   models.Product
		ProductID int
		Price     int
		SubTotal  int
	}

	c.Bind(&productOrder)

	temp_product := models.ProductOrder{OrderID: productOrder.OrderID, Quantity: productOrder.Quantity, Price: productOrder.Price, SubTotal: productOrder.SubTotal, ProductID: productOrder.ProductID, Product: productOrder.Product}

	result := initializers.DB.Create(&temp_product)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"Product Quantity": productOrder.Quantity,
		"Product Price":    productOrder.Price,
	})
}

func RetrieveAllProductOrders(c *gin.Context) {
	//var records []models.Product
	var productOrders []models.ProductOrder
	initializers.DB.Find(&productOrders)

	c.JSON(200, gin.H{
		"Products": productOrders,
	})
}

func RetrieveProductOrderByIndex(c *gin.Context) {
	var productOrder models.Product
	index := c.Param("index")
	initializers.DB.Find(&productOrder, index)
	fmt.Println(index)
	c.JSON(200, gin.H{
		"ProductOrder": productOrder,
	})
}

func UpdateProductOrderByIndex(c *gin.Context) {
	var productOrder models.ProductOrder
	var userInput models.ProductOrder
	c.Bind(&userInput)

	index := c.Param("index")
	initializers.DB.First(&productOrder, index)
	if userInput.OrderID != 0 {
		initializers.DB.Model(&productOrder).Update("OrderD", userInput.OrderID)
	}

	if userInput.ProductID != 0 {
		initializers.DB.Model(&productOrder).Update("ProductID", userInput.ProductID)
	}

	if userInput.Quantity != 0 {
		initializers.DB.Model(&productOrder).Update("Quantity", userInput.Quantity)
	}

	if userInput.Price != 0 {
		initializers.DB.Model(&productOrder).Update("Price", userInput.Price)
	}

	if userInput.SubTotal != 0 {
		initializers.DB.Model(&productOrder).Update("SubTotal", userInput.SubTotal)
	}

	c.JSON(200, gin.H{
		"ProductOrder": productOrder,
	})
}

func DeleteProductOrderByIndex(c *gin.Context) {
	index := c.Param("index")
	initializers.DB.Delete(&models.ProductOrder{}, index)
	fmt.Println(index)
	c.Status(200)
}
