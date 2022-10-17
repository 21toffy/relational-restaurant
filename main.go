package main

import (
	// "github.com/21toffy/relational-restaurant/controllers"
	"fmt"
	"log"

	"github.com/21toffy/relational-restaurant/database"
	"github.com/21toffy/relational-restaurant/models"
	"github.com/21toffy/relational-restaurant/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.DB, err = gorm.Open(postgres.Open(database.DbURL(database.BuildDBConfig())))

	if err != nil {
		fmt.Println("Status:", err)
	}
	// defer database.DB.Close()
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Food{})
	database.DB.AutoMigrate(&models.Table{})
	database.DB.AutoMigrate(&models.Menu{})
	database.DB.AutoMigrate(&models.OrderItem{})
	database.DB.AutoMigrate(&models.Order{})

	router := gin.New()
	router.Use(gin.Logger())

	routes.FoodRoutes(router)
	routes.UserRoutes(router)
	routes.TableRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderItemRoutes(router)
	routes.OrderRoutes(router)

	router.Run(":8082")

}
