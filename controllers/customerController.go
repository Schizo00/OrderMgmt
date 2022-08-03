package controllers

import (
	"OrderMgmt/initializers"
	"OrderMgmt/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateCustomer(c *gin.Context) {

	var customer struct {
		Name   string `json:"Name"`
		Email  string `json:"Email"`
		Orders []models.Order
	}

	c.Bind(&customer)

	temp_order := models.Customer{Name: customer.Name, Email: customer.Email, Orders: customer.Orders}

	result := initializers.DB.Create(&temp_order)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"Customer ID": customer.Name,
		"Email":       customer.Email,
	})
}

func RetrieveAllCustomers(c *gin.Context) {
	//var records []models.Product
	var customers []models.Customer
	initializers.DB.Find(&customers)

	c.JSON(200, customers)
}

func RetrieveCustomerByIndex(c *gin.Context) {
	var customer models.Customer
	index := c.Param("index")
	initializers.DB.Find(&customer, index)
	fmt.Println(index)
	c.JSON(200, gin.H{
		"Customer": customer,
	})
}

func UpdateCustomerByIndex(c *gin.Context) {
	var customer models.Customer

	var userInput struct {
		Name   string `json:"Name"`
		Email  string `json:"Email"`
		Orders []models.Order
	}
	c.Bind(&userInput)

	index := c.Param("index")
	initializers.DB.First(&customer, index)
	if userInput.Name != "" {
		initializers.DB.Model(&customer).Update(models.Customer{Name: userInput.Name})
	}

	if userInput.Email != "" {
		initializers.DB.Model(&customer).Update(models.Customer{Email: userInput.Email})
	}

	c.JSON(200, gin.H{
		"Customer": customer,
	})
}

func DeleteCustomerByIndex(c *gin.Context) {
	index := c.Param("index")
	initializers.DB.Delete(&models.Customer{}, index)
	fmt.Println(index)
	c.Status(200)
}
