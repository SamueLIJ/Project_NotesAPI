package route

import (
	"NotesAPI/controller"
	"NotesAPI/middleware"

	"github.com/labstack/echo"
)

func NewNotes(e *echo.Echo) {
	e.GET("/notes", controller.GetAllNotesController, middleware.AuthJWT)
	e.POST("/notes", controller.CreateNoteController, middleware.AuthJWT)
	e.GET("/notes/:id", controller.GetNoteByIDController, middleware.AuthJWT)
	e.DELETE("/notes/:id", controller.DeleteNoteByIDController, middleware.AuthJWT)
	e.PUT("/notes/:id", controller.UpdateNoteByIDController, middleware.AuthJWT)
}
