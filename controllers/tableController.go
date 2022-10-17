package controller

import (
	"fmt"
	"net/http"

	"github.com/21toffy/relational-restaurant/helpers"
	"github.com/21toffy/relational-restaurant/models"
	"github.com/gin-gonic/gin"
)

type createTableStruct struct {
	Number_of_guests *int64 `json:"number_of_guests" binding:"required"`
	Table_number     *int64 `json:"table_number" binding:"required"`
}

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tables []models.Table
		err := models.GetAllTables(&tables)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, gin.H{"data": tables})
		}
	}
}

func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input createTableStruct
		c.BindJSON(&input)
		table := models.Table{Table_id: helpers.GenerateUUID(), Number_of_guests: input.Number_of_guests, Table_number: input.Table_number}
		err := models.CreateTable(&table)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Table created successfully", "data": input})
			return
		}

	}
}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
