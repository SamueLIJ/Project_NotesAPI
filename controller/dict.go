package controller

import (
	"NotesAPI/lib"
	"net/http"

	"github.com/labstack/echo"
)

func GetWordDefinitionController(c echo.Context) error {
	words := c.Param("word")
	
	result := lib.GetWord("en", words)
	
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Get word",
		"dictionary": result,
	})
}
