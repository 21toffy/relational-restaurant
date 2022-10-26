package controller

import (
	"net/http"
	"strconv"

	"github.com/21toffy/relational-restaurant/helpers"
	"github.com/21toffy/relational-restaurant/models"
	"github.com/gin-gonic/gin"
)

type MenuStruct struct {
	Name     string `json:"name" binding:"required,gte=1"`
	Category string `json:"category" binding:"required,min=3"`
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input MenuStruct
		c.BindJSON(&input)

		if input.Category == "" || input.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "category and name input can not be empty"})
			return
		}
		menu := models.Menu{Menu_id: helpers.GenerateUUID(), Name: input.Name, Category: input.Category}
		err := models.CreateMenu(&menu)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Menu created successfully", "data": input})
			return
		}

	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		menuId, _ := strconv.Atoi(c.Param("menu_id"))
		var menu models.Menu
		menu, err := models.GetMenuByID(menuId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, menu)
		return

	}
}

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var menus []models.Menu
		err := models.GetAllMenus(&menus)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, gin.H{"data": menus})
		}
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
