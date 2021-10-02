package route

import (
	"NotesAPI/controller"

	"github.com/labstack/echo"
)

func NewLabel(e *echo.Echo) {
	e.GET("/labels", controller.GetAllLabelsController)
	e.POST("/labels", controller.CreateLabelController)
	e.GET("/labels/:id", controller.GetLabelByIDController)
	e.DELETE("/labels/:id", controller.DeleteLabelByIDController)
	e.PUT("/labels/:id", controller.UpdateLabelByIDController)
}
