package controller

import (
	"NotesAPI/lib/database"
	"NotesAPI/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllNotesController(c echo.Context) error {
	notes := database.GetNotes()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Get all notes",
		"data":    notes,
	})
}

func GetNoteByIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	notes, err := database.GetNotesByID(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Get note by ID",
		"data":    notes,
	})
}

func DeleteNoteByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeleteNoteByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Note Successfully Deleted",
	})
}

func UpdateNoteByIDController(c echo.Context) error {
	id := c.Param("id")

	var note model.Note
	if err := c.Bind(&note); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "error update note",
			"error":   err.Error(),
		})
	}
	database.UpdateNoteByID(id, note)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Note Successfully Updated",
	})
}

func CreateNoteController(c echo.Context) error {
	var newNote model.Note
	if err := c.Bind(&newNote); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "error create note",
			"error":   err.Error(),
		})
	}

	newNote = database.CreateNote(newNote)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Note Successfully Created",
		"noteID":  newNote.ID,
	})
}
