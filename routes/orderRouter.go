package routes

import (
	controller "github.com/21toffy/relational-restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(incommingRoutes *gin.Engine) {

	public := incommingRoutes.Group("/api")
	// private := incommingRoutes.Group("/api").Use(middleware.JwtAuthMiddleware())

	public.GET("/orders", controller.GetOrders())
	public.GET("/orders/:order_id", controller.GetOrder())
	public.POST("/orders", controller.CreateOrder())
	public.PATCH("/orders/:order_id", controller.UpdateOrder())

}
