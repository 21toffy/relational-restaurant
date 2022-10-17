package routes

import (
	controller "github.com/21toffy/relational-restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func TableRoutes(incommingRoutes *gin.Engine) {

	public := incommingRoutes.Group("/api")
	// private := incommingRoutes.Group("/api").Use(middleware.JwtAuthMiddleware())

	public.GET("/tables", controller.GetTables())
	public.GET("/tables/:table_id", controller.GetTable())
	public.POST("/tables", controller.CreateTable())
	public.PATCH("/tables/:table_id", controller.UpdateTable())

}
