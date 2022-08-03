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
	c.JSON(200, gin.H{
		"ProductID": order,
	})
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

	c.JSON(200, gin.H{
		"UserOrder": userInput,
	})

	c.JSON(200, gin.H{
		"Order": order,
	})

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

// func CalculateTotalPrice(prodOrderPrice int, index int, c *gin.Context) int {
// 	var temp_all_orders []models.ProductOrder
// 	//initializers.DB.Find(&temp_all_orders, index)
// 	initializers.DB.Raw("SELECT * FROM product_orders WHERE order_id = ?", index).Scan(&temp_all_orders)
// 	temp_total_price := prodOrderPrice

// 	for i := 0; i < len(temp_all_orders); i++ {
// 		temp_total_price = temp_total_price + temp_all_orders[i].SubTotal
// 	}

// 	c.JSON(200, gin.H{
// 		"NO OF ELEMENTS": len(temp_all_orders),
// 	})

// 	return temp_total_price
// }

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

		initializers.DB.Model(&orders[i]).Update("Total", temp_total)

	}

}
