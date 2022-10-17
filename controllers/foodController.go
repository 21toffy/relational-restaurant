package controller

import (
	"fmt"
	"math"
	"net/http"

	"github.com/21toffy/relational-restaurant/helpers"
	"github.com/21toffy/relational-restaurant/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {

		var food []models.Food
		err := models.GetAllFoods(&food)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, gin.H{"data": food})
		}
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

type FoodStruct struct {
	Name        *string  `json:"name" binding:"required"`
	Price       *float64 `json:"price" binding:"required"`
	Description *string  `json:"description" binding:"required"`
	Food_image  *string  `json:"food_image" binding:"required"`
	Created_by  models.User
	Created_at  string
	Updated_at  string
	Food_id     string
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input FoodStruct
		c.BindJSON(&input)
		user_id, err := helpers.ExtractTokenID(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		u, err := models.GetCurrentID(user_id)
		fmt.Println(u, 666)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		food := models.Food{
			Name:        input.Name,
			Price:       input.Price,
			Description: input.Description,
			Food_image:  input.Food_image,
			User:        u,
			Created_at:  helpers.GetCurrentTime(),
			Updated_at:  helpers.GetCurrentTime(),
			Food_id:     helpers.GenerateUUID(),
			UserId:      int(u.Id),
		}

		errors := models.CreateFood(&food)

		if errors != nil {
			fmt.Println(errors.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": errors.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"name": food.Name, "price": food.Price, "created_by": food.User})
			return
		}
	}
}

// func CreateFood() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 	}
// }

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
