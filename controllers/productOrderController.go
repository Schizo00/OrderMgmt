package controllers

import (
	"OrderMgmt/initializers"
	"OrderMgmt/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateProductOrder(c *gin.Context) {

	var productOrder struct {
		OrderID   int `json:"OrderID"`
		Quantity  int `json:"Quantity"`
		Order     models.Order
		Product   models.Product
		ProductID int `json:"ProductId"`
		Price     int `json:"Price"`
		SubTotal  int `json:"SubTotal"`
	}

	c.Bind(&productOrder)
	fmt.Println("OrderID", productOrder.OrderID)
	var temp_product models.Product
	initializers.DB.First(&temp_product, productOrder.ProductID)

	temp_productOrder := models.ProductOrder{
		OrderID:   productOrder.OrderID,
		Quantity:  productOrder.Quantity,
		ProductID: productOrder.ProductID,
		Product:   productOrder.Product,
		Price:     temp_product.Price,
		SubTotal:  temp_product.Price * productOrder.Quantity,
	}
	result := initializers.DB.Create(&temp_productOrder)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"Product Quantity": productOrder.Quantity,
		"Product Price":    productOrder.Price,
	})

	refreshTotals()
}

func RetrieveAllProductOrders(c *gin.Context) {

	refreshTotals()

	//var records []models.Product
	var productOrders []models.ProductOrder
	initializers.DB.Find(&productOrders)

	c.JSON(200, productOrders)
}

func RetrieveProductOrderByIndex(c *gin.Context) {

	refreshTotals()

	var productOrder models.ProductOrder
	index := c.Param("index")
	initializers.DB.Find(&productOrder, index)
	fmt.Println(index)
	c.JSON(200, productOrder)
}

func UpdateProductOrderByIndex(c *gin.Context) {
	var productOrder models.ProductOrder

	var userInput struct {
		OrderID   int `json:"OrderID"`
		Quantity  int `json:"Quantity"`
		Order     models.Order
		Product   models.Product
		ProductID int `json:"ProductId"`
		Price     int `json:"Price"`
		SubTotal  int `json:"SubTotal"`
	}

	c.Bind(&userInput)

	index := c.Param("index")
	initializers.DB.First(&productOrder, index)

	fmt.Println("INDEX: ", index)
	if userInput.OrderID != 0 {
		initializers.DB.Model(&productOrder).Update(models.ProductOrder{OrderID: userInput.OrderID})
	}

	if userInput.ProductID != 0 {
		initializers.DB.Model(&productOrder).Update(models.ProductOrder{ProductID: userInput.ProductID})
	}

	if userInput.Quantity != 0 {
		initializers.DB.Model(&productOrder).Update(models.ProductOrder{Quantity: userInput.Quantity})
	}

	var temp_product models.Product
	initializers.DB.Find(&temp_product, productOrder.ProductID)

	initializers.DB.Model(&productOrder).Update(models.ProductOrder{Price: temp_product.Price})
	initializers.DB.Model(&productOrder).Update(models.ProductOrder{SubTotal: temp_product.Price * productOrder.Quantity})

	c.JSON(200, gin.H{
		"Order": 1,
	})

	fmt.Println("INSIDE UPDATE FUNC")
	refreshTotals()

}

func DeleteProductOrderByIndex(c *gin.Context) {
	index := c.Param("index")
	var productOrder models.ProductOrder
	initializers.DB.First(&productOrder, index)
	initializers.DB.Delete(&models.ProductOrder{}, index)
	c.JSON(200, gin.H{
		"Order": 1,
	})
	fmt.Println(index)
	c.Status(200)

	refreshTotals()
}
