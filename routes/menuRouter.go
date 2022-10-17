package routes

import (
	controller "github.com/21toffy/relational-restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func MenuRoutes(incommingRoutes *gin.Engine) {
	public := incommingRoutes.Group("/api")
	// private := incommingRoutes.Group("/api").Use(middleware.JwtAuthMiddleware())
	public.GET("/menus", controller.GetMenus())
	public.GET("/menus/:menu_id", controller.GetMenu())
	public.POST("/menus", controller.CreateMenu())
	public.PATCH("/menus/:menu_id", controller.UpdateMenu())

}
