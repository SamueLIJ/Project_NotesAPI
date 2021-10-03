package controller

import (
	"NotesAPI/lib/database"
	"NotesAPI/model"
	"net/http"
	"strconv"

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
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	labels, err := database.GetLabelsByID(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "login success",
		"message": "NotebyID",
		"data":    labels,
	})
}

func DeleteLabelByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeleteLabelByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Label Successfully Deleted",
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
		"message": "Label Successfully Updated",
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
