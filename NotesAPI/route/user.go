package route

import (
	"NotesAPI/controller"

	"github.com/labstack/echo"
)

func NewUser(e *echo.Echo) {
	e.GET("/users", controller.GetAllUsersController)
	e.POST("/users/login", controller.LoginUserController)
	e.POST("/users", controller.CreateUserController)
	e.GET("/users/:id", controller.GetUserDetailController)
	e.DELETE("/users/:id", controller.DeleteUserByIDController)
	e.PUT("/users/:id", controller.UpdateUserByIDController)
}