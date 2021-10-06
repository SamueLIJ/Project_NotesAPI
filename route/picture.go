package route

import (
	"NotesAPI/controller"
	"NotesAPI/middleware"

	"github.com/labstack/echo"
)

func NewPictures(e *echo.Echo) {
	e.GET("/notes/pictures", controller.GetAllPicturesController, middleware.AuthJWT)
	e.POST("/notes/pictures", controller.CreatePictureController, middleware.AuthJWT)
	e.GET("/notes/pictures/:id", controller.GetPictureByIDController, middleware.AuthJWT)
	e.DELETE("/notes/pictures/:id", controller.DeletePictureByIDController, middleware.AuthJWT)
	e.PUT("/notes/pictures/:id", controller.UpdatePictureByIDController,middleware.AuthJWT)
}
