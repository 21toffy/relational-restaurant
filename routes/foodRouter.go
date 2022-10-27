package routes

import (
	controller "github.com/21toffy/relational-restaurant/controllers"
	"github.com/21toffy/relational-restaurant/middleware"
	"github.com/gin-gonic/gin"
)

func FoodRoutes(incommingRoutes *gin.Engine) {
	public := incommingRoutes.Group("/api")

	private := incommingRoutes.Group("/api").Use(middleware.JwtAuthMiddleware())

	public.GET("/foods", controller.GetFoods())
	public.GET("/foods/:food_id", controller.GetFood())
	private.POST("/food", controller.CreateFood())
	public.PATCH("/foods/:food_id", controller.UpdateFood())
	public.DELETE("/foods/:food_id", controller.DeleteFood())

}
