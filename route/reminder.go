package route

import (
	"NotesAPI/controller"
	"NotesAPI/middleware"

	"github.com/labstack/echo"
)

func NewReminders(e *echo.Echo) {
	e.GET("/notes/reminders", controller.GetAllRemindersController, middleware.AuthJWT)
	e.POST("/notes/reminders", controller.CreateReminderController,middleware.AuthJWT)
	e.GET("/notes/reminders/:id", controller.GetReminderByIDController,middleware.AuthJWT)
	e.DELETE("/notes/reminders/:id", controller.DeleteReminderByIDController,middleware.AuthJWT)
	e.PUT("/notes/reminders/:id", controller.UpdateReminderByIDController,middleware.AuthJWT)
}
