package controller

import (
	// "fmt"

	"fmt"
	"net/http"

	// "strconv"

	// "github.com/21toffy/relational-restaurant/models"
	// "net/http"

	// "github.com/21toffy/relational-restaurant/models"

	"github.com/21toffy/relational-restaurant/helpers"
	"github.com/21toffy/relational-restaurant/models"
	"github.com/gin-gonic/gin"
)

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

		var user []models.User
		err := models.GetAllUsers(&user)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, gin.H{"data": user})
		}
	}
}

// func GetUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var user models.User

// 		// if err := c.ShouldBindJSON(&user); err != nil {
// 		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		// 	return
// 		// }

// 		err := Models.GetUserByID(&user, id)

// 		fmt.Println(c.Param("user_id"), 9898)
// 		// c:=c.Param("user_id")
// 		id, _ := strconv.Atoi(c.Param("user_id"))
// 		c.JSON(http.StatusOK, gin.H{"1": id, "2": c.Param("user_id")})
// 	}
// }

type signUpStruct struct {
	Name     string
	Email    string
	Phone    string
	Address  string
	Password string
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input LoginInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		u := models.User{}

		u.Email = input.Email
		u.Password = input.Password

		token, err := models.LoginCheck(u.Email, u.Password)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "username or password is incorrect."})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})

	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input signUpStruct
		c.BindJSON(&input)
		user := models.User{Uid: helpers.GenerateUUID(), Name: input.Name, Email: input.Email, Phone: input.Phone, Address: input.Address, Password: helpers.HashPassword(input.Password)}
		err := models.CreateUser(&user)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			input.Password = ""
			c.JSON(http.StatusOK, gin.H{"message": "user fetched successfully", "data": input})
			return
		}
	}
}
