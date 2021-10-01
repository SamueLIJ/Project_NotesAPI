package route

import (
	"NotesAPI/controller"

	"github.com/labstack/echo"
)

func NewLabel(e *echo.Echo) {
	e.GET("/notes/labels", controller.GetAllLabelsController)
	e.POST("/notes/labels", controller.CreateLabelController)
	e.GET("/notes/labels/:id", controller.GetLabelByIDController)
	e.DELETE("/notes/labels/:id", controller.DeleteLabelByIDController)
	e.PUT("/notes/labels/:id", controller.UpdateLabelByIDController)
}
