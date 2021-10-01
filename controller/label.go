package controller

import (
	"NotesAPI/lib/database"
	"NotesAPI/model"
	"net/http"
	"github.com/labstack/echo"
)

func GetAllLabelsController(c echo.Context) error {
	labels := database.GetLabels()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetAllLabelsController",
		"data":    labels,
	})
}

func GetLabelByIDController(c echo.Context) error {
	id := c.Param("id")
	label := database.GetLabelsByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetLabelByIDController",
		"data":    label,
	})
}

func DeleteLabelByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeleteLabelByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "DeleteLabelByIDController",
	})
}

func UpdateLabelByIDController(c echo.Context) error {
	id := c.Param("id")

	var label model.Label
	if err := c.Bind(&label); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateLabelController",
			"error":   err.Error(),
		})
	}
	database.UpdateLabelByID(id, label)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetLabelByIDController",
		"data":    label,
	})
}

func CreateLabelController(c echo.Context) error {
	var newLabel model.Label
	if err := c.Bind(&newLabel); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateLabelController",
			"error":   err.Error(),
		})
	}

	newLabel = database.CreateLabel(newLabel)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreateLabelController",
		"data":    newLabel,
	})
}
