package controller

import (
	"NotesAPI/lib/database"
	"NotesAPI/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllRemindersController(c echo.Context) error {
	reminders := database.GetReminders()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetAllRemindersController",
		"data":    reminders,
	})
}

func GetReminderByIDController(c echo.Context) error {
	id := c.Param("id")
	reminder := database.GetReminderByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetReminderByIDController",
		"data":    reminder.ID,
	})
}

func DeleteReminderByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeleteReminderByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "DeleteReminderByIDController",
	})
}

func UpdateReminderByIDController(c echo.Context) error {
	id := c.Param("id")

	var reminder model.Reminder
	if err := c.Bind(&reminder); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateReminderController",
			"error":   err.Error(),
		})
	}
	database.UpdateReminderByID(id, reminder)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetReminderByIDController",
		"data":    reminder,
	})
}

func CreateReminderController(c echo.Context) error {
	var newReminder model.Reminder
	if err := c.Bind(&newReminder); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateReminderController",
			"error":   err.Error(),
		})
	}

	newReminder = database.CreateReminder(newReminder)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreateReminderController",
		"data":    newReminder,
	})
}
