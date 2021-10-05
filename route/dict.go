package route

import (
	"NotesAPI/controller"

	"github.com/labstack/echo"
)

func Dict(e *echo.Echo) {
	e.GET("/dict/:word", controller.GetWordDefinitionController)
}
