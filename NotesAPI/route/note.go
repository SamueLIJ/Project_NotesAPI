package route

import (
	"NotesAPI/controller"

	"github.com/labstack/echo"
)

func NewNotes(e *echo.Echo) {
	e.GET("/notes", controller.GetAllNotesController)
	e.POST("/notes", controller.CreateNoteController)
	e.GET("/notes/:id", controller.GetNoteByIDController)
	e.DELETE("/notes/:id", controller.DeleteNoteByIDController)
	e.PUT("/notes/:id", controller.UpdateNoteByIDController)
}
