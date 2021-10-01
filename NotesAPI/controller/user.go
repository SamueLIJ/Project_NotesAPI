package controller

import (
	"NotesAPI/lib/database"
	"NotesAPI/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllUsersController(c echo.Context) error {
	users := database.GetUsers()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetAllUsersController",
		"data":    users,
	})
}

func GetUserByIDController(c echo.Context) error {
	id := c.Param("id")
	user := database.GetUserByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetUserByIDController",
		"data":    user,
	})
}

func DeleteUserByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeleteUserByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "DeleteUserByIDController",
	})
}

func UpdateUserByIDController(c echo.Context) error {
	id := c.Param("id")

	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateUserController",
			"error":   err.Error(),
		})
	}
	database.UpdateUserByID(id, user)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetUserByIDController",
		"data":    user,
	})
}

func CreateUserController(c echo.Context) error {
	var newUser model.User
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateUserController",
			"error":   err.Error(),
		})
	}

	newUser = database.CreateUser(newUser)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreateUserController",
		"data":    newUser,
	})

}

// func LoginUserController(c echo.Context) error {
// 	var user model.User

// 	isValid := database.IsValid(user.Email, user.Name)
// 	if !isValid {
// 		return c.String(http.StatusBadRequest, "Email atau Password Salah")
// 	}

// 	claims := jwt.MapClaims{}

// 	claims["userID"] = user.Email
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString([]byte("legal"))

// 	if err != nil {
// 		return c.String(http.StatusInternalServerError, err.Error())
// 	}

// 	return c.String(http.StatusOK, tokenString)
// }

func LoginUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	users, err := database.LoginUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "login success",
		"user":   users,
	})

}

func GetUserDetailController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	users, err := database.GetDetailUser(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "login success",
		"user":   users,
	})

}
