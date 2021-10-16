package controller

import (
	"NotesAPI/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func Initecho() *echo.Echo {
	config.InitDB()
	e := echo.New()

	return e
}

func TestGetUserbyIDController(t *testing.T) {
	var testcases = []struct {
		name            string
		path            string
		expectedStatus  int
		expectbodystart string
	}{
		{
			name:            "Success ",
			path:            "/users",
			expectbodystart: "{\"status\":\"success\",\"users\":[",
			expectedStatus:  http.StatusOK,
		},
	}

	e := Initecho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testcase := range testcases {
		c.SetPath(testcase.path)

		if assert.NoError(t, GetAllUsersController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			rec.Body.String()
			// body := rec.Body.String()
			// assert.True(t, strings.HasPrefix(body, testcase.expectbodystart))
		}

	}
}

func TestCreateUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/users", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)

	err := CreateUserController(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}

func TestUpdateUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "/users/:id", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)
	c.Set("user", jwt.MapClaims{
		"userId": float64(1),
	})
	err := UpdateUserByIDController(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}
