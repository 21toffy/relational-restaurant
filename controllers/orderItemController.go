package controller

import (
	"net/http"
	"time"

	"github.com/21toffy/relational-restaurant/models"
	"github.com/gin-gonic/gin"
)

type OrderItemCreate struct {
	Size       string
	Quantity   string
	Unit_price float64
	FoodId     string
	OrderId    string
}

type OrderItemFood struct {
	Name       string
	Price      float64
	Food_image string
}

type OrderItemOrder struct {
	Number_of_guests int64
	Table_number     int64
	Order_Date       time.Time
}

type OrderItemList struct {
	Order_Item_id string
	Size          string
	Quantity      string
	Unit_price    float64
	Food          OrderItemFood
	Order         OrderItemOrder
}

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		var orderItems []models.OrderItem
		err := models.GetAllOrderItem(&orderItems)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"data": orderItems})
			return
		}

	}
}

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		order_id := c.Param("order_id")
		var orderItems []models.OrderItem
		err := models.GetOrderItemsByOrder(order_id, &orderItems)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"data": orderItems})
			return
		}
	}
}
func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// func ItemsByOrder(id string) (orderItems []primitive.M, err error) {

// }
