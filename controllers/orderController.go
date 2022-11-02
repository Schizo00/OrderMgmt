package controllers

import (
	"OrderMgmt/initializers"
	"OrderMgmt/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {

	var order struct {
		CustID        int `json:"CustID"`
		ProductOrders []models.ProductOrder
		Total         int `json:"Total"`
	}

	c.Bind(&order)
	//
	temp_order := models.Order{CustID: order.CustID, ProductOrders: order.ProductOrders}

	result := initializers.DB.Create(&temp_order)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"Customer ID": order.CustID,
		"Total":       order.Total,
	})

	refreshTotals()
}

func RetrieveAllOrders(c *gin.Context) {

	refreshTotals()

	//var records []models.Product
	var orders []models.Order
	initializers.DB.Find(&orders)

	c.JSON(200, orders)

	refreshTotals()
}

func RetrieveOrderByIndex(c *gin.Context) {

	refreshTotals()

	var order models.Order
	index := c.Param("index")
	initializers.DB.Find(&order, index)
	fmt.Println(index)
	c.JSON(200, order)

	refreshTotals()
}

func UpdateOrderByIndex(c *gin.Context) {
	var order models.Order
	var userInput struct {
		CustID        int `json:"CustID"`
		ProductOrders []models.ProductOrder
		Total         int `json:"Total"`
	}
	c.Bind(&userInput)

	fmt.Println(userInput, "User null")

	index := c.Param("index")
	initializers.DB.First(&order, index)
	if userInput.CustID != 0 {
		initializers.DB.Model(&order).Update(models.Order{CustID: userInput.CustID})
	}

	refreshTotals()

}

func DeleteOrderByIndex(c *gin.Context) {
	index := c.Param("index")
	initializers.DB.Delete(&models.Order{}, index)
	fmt.Println(index)
	c.JSON(200, gin.H{
		"In Func": 100,
	})

	refreshTotals()

}

func refreshTotals() {
	var prod_orders []models.ProductOrder
	initializers.DB.Find(&prod_orders)

	var orders []models.Order
	initializers.DB.Find(&orders)

	for i := 0; i < len(orders); i++ {
		initializers.DB.Model(&orders[i]).Update("Total", 0)
	}

	for i := 0; i < len(prod_orders); i++ {
		var order models.Order
		initializers.DB.First(&order, prod_orders[i].OrderID)

		temp_total := order.Total + prod_orders[i].SubTotal
		fmt.Println("SUBTOTAL: ", temp_total)
		fmt.Println("PROD ORDER: ", prod_orders[i].OrderID)

		initializers.DB.Model(&order).Update("Total", temp_total)
		fmt.Println("ORDER ID: ", order.OrderID)

	}

	fmt.Println("INSIDE REFRESH TOTALS FUNC")

}

func RefreshTotal(c *gin.Context) {
	var prod_orders []models.ProductOrder
	initializers.DB.Find(&prod_orders)

	var orders []models.Order
	initializers.DB.Find(&orders)

	for i := 0; i < len(orders); i++ {
		initializers.DB.Model(&orders[i]).Update("Total", 0)
	}

	for i := 0; i < len(prod_orders); i++ {
		var order models.Order
		initializers.DB.First(&order, prod_orders[i].OrderID)

		temp_total := order.Total + prod_orders[i].SubTotal
		fmt.Println("SUBTOTAL: ", temp_total)

		initializers.DB.Model(&orders[i]).Update("Total", temp_total)

	}

	fmt.Println("INSIDE REFRESH TOTALS FUNC")

	c.JSON(200, gin.H{
		"HI": "Done",
	})

}
