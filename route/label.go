package route

import (
	"NotesAPI/controller"
	"NotesAPI/middleware"

	"github.com/labstack/echo"
)

func NewLabel(e *echo.Echo) {
	e.GET("/labels", controller.GetAllLabelsController,middleware.AuthJWT)
	e.POST("/labels", controller.CreateLabelController, middleware.AuthJWT)
	e.GET("/labels/:id", controller.GetLabelByIDController, middleware.AuthJWT)
	e.DELETE("/labels/:id", controller.DeleteLabelByIDController, middleware.AuthJWT)
	e.PUT("/labels/:id", controller.UpdateLabelByIDController, middleware.AuthJWT)
}
