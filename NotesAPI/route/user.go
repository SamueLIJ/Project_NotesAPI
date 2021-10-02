package route

import (
	"NotesAPI/controller"
	"fmt"

	"github.com/labstack/echo"
)

func NewUser(e *echo.Echo) {
	e.Use(func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Println("hello")
			return hf(c)
		}
	})

	e.GET("/users", controller.GetAllUsersController)
	e.POST("/users/login", controller.LoginUserController)
	e.POST("/users", controller.CreateUserController)
	e.GET("/users/:id", controller.GetUserDetailController)
	e.DELETE("/users/:id", controller.DeleteUserByIDController)
	e.PUT("/users/:id", controller.UpdateUserByIDController)
}
