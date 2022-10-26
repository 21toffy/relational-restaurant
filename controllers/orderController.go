package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/21toffy/relational-restaurant/helpers"
	"github.com/21toffy/relational-restaurant/models"
	"github.com/gin-gonic/gin"
)

type OrderTableDisplay struct {
	Number_of_guests *int64 `json:"number_of_guests" binding:"required"`
	Table_number     *int64 `json:"table_number" binding:"required"`
}

type CreateOrderStruct struct {
	Table_id   string    `json:"table_id" validate:"required"`
	Order_Date time.Time `json:"order_date" validate:"required"`
}

type OrderDisplay struct {
	Table      OrderTableDisplay
	Order_Date time.Time `json:"order_date" validate:"required"`
}

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		var order []models.Order
		err := models.GetAllOrders(&order)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"data": order})
			return
		}
	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

		orderId, _ := strconv.Atoi(c.Param("order_id"))
		orderData, err := models.GetOrderByID(orderId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, orderData)
		return

	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input CreateOrderStruct
		c.BindJSON(&input)
		if input.Table_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "table id can not be empty"})
			return
		}
		order := models.Order{Order_id: helpers.GenerateUUID(), Table_id: &input.Table_id}
		order.Order_Date, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		err := models.CreateOrder(&order)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"data": order})
			return
		}
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
