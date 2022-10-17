package routes

import (
	controller "github.com/21toffy/relational-restaurant/controllers"
	"github.com/21toffy/relational-restaurant/middleware"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incommingRoutes *gin.Engine) {
	public := incommingRoutes.Group("/api")

	private := incommingRoutes.Group("/api").Use(middleware.JwtAuthMiddleware())

	public.GET("/orderItems", controller.GetOrderItems())
	public.POST("/orderItems", controller.CreateOrderItem())
	public.GET("/orderItems/:orderItem_id", controller.GetOrderItem())
	public.GET("/orderItems-order/order_id", controller.GetOrderItemsByOrder())
	private.PATCH("/orderItems/:orderItem_id", controller.UpdateOrderItem())

}
