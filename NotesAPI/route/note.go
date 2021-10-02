package route

import (
	"NotesAPI/controller"
	"NotesAPI/middleware"

	"github.com/labstack/echo"
)

func NewNotes(e *echo.Echo) {
	e.GET("/notes", controller.GetAllNotesController)
	e.POST("/notes", controller.CreateNoteController)
	e.GET("/notes/:id", controller.GetNoteByIDController, middleware.AuthJWT)
	e.DELETE("/notes/:id", controller.DeleteNoteByIDController)
	e.PUT("/notes/:id", controller.UpdateNoteByIDController)
}
