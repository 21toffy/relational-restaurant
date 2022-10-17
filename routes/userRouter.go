package routes

import (
	controller "github.com/21toffy/relational-restaurant/controllers"
	"github.com/21toffy/relational-restaurant/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incommingRoutes *gin.Engine) {
	public := incommingRoutes.Group("/api")
	private := incommingRoutes.Group("/api").Use(middleware.JwtAuthMiddleware())

	private.GET("/users", controller.GetUsers())
	public.POST("/users/signup", controller.SignUp())
	public.POST("/users/login", controller.Login())
}
