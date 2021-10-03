package route

import (
	"NotesAPI/controller"

	"github.com/labstack/echo"
)

func NewPictures(e *echo.Echo) {
	e.GET("/notes/pictures", controller.GetAllPicturesController)
	e.POST("/notes/pictures", controller.CreatePictureController)
	e.GET("/notes/pictures/:id", controller.GetPictureByIDController)
	e.DELETE("/notes/pictures/:id", controller.DeletePictureByIDController)
	e.PUT("/notes/pictures/:id", controller.UpdatePictureByIDController)
}
