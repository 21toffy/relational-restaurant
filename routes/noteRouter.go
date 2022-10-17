package routes

import (
	controller "github.com/21toffy/relational-restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func NoteRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/notes", controller.GetNotess())
	incommingRoutes.GET("/notes/:note_id", controller.GetNotes())
	incommingRoutes.POST("/notes", controller.CreateNotes())
	incommingRoutes.PATCH("/notes/:note_id", controller.UpdateNotes())

}
