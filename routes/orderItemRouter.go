package routes

import (
	controller "github.com/21toffy/relational-restaurant/controllers"
	"github.com/21toffy/relational-restaurant/middleware"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incommingRoutes *gin.Engine) {
	public := incommingRoutes.Group("/api")

	private := incommingRoutes.Group("/api").Use(middleware.JwtAuthMiddleware())

	public.GET("/order-items", controller.GetOrderItems())
	public.POST("/order-items", controller.CreateOrderItem())
	public.GET("/order-items/:order-item-id", controller.GetOrderItem())
	public.GET("/order-items-order/:order_id", controller.GetOrderItemsByOrder())
	private.PATCH("/orderItems/:orderItem_id", controller.UpdateOrderItem())

}
