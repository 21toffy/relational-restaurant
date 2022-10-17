package controller

import (
	"net/http"
	"time"

	"github.com/21toffy/relational-restaurant/helpers"
	"github.com/21toffy/relational-restaurant/models"
	"github.com/gin-gonic/gin"
)

type OrderItemCreate struct {
	Size       string
	Quantity   int64
	Unit_price float64
	FoodId     int
	OrderId    int
}

type OrderItemFood struct {
	Name       string
	Price      float64
	Food_image string
}

type OrderItemOrder struct {
	Order_id   string
	Order_Date time.Time
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
		// order_id, _ := strconv.Atoi(c.Param("order_id"))
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
		var input OrderItemCreate

		c.BindJSON(&input)
		if input.Size == "" || input.Quantity <= 0 || input.Unit_price <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "size, quantity, unit price fields can not be empty"})
			return
		}

		order, err := models.GetOrderByID(input.OrderId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		food, err := models.GetFoodByID(input.FoodId)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		orderItem := models.OrderItem{
			Order_Item_id: helpers.GenerateUUID(),
			Size:          input.Size,
			Quantity:      input.Quantity,
			Unit_price:    input.Unit_price,

			FoodId:  input.FoodId,
			OrderId: input.OrderId,
			Order:   order,
			Food:    food,
		}

		errors := models.CreateOrderItem(&orderItem)

		if errors != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errors.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"data": orderItem})
			return
		}
	}
}
func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// func ItemsByOrder(id string) (orderItems []primitive.M, err error) {

// }
