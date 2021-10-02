package route

import (
	"NotesAPI/controller"

	"github.com/labstack/echo"
)

func NewReminders(e *echo.Echo) {
	e.GET("/notes/reminders", controller.GetAllRemindersController)
	e.POST("/notes/reminders", controller.CreateReminderController)
	e.GET("/notes/reminders/:id", controller.GetReminderByIDController)
	e.DELETE("/notes/reminders/:id", controller.DeleteReminderByIDController)
	e.PUT("/notes/reminders/:id", controller.UpdateReminderByIDController)
}
