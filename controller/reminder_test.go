package controller

import (
	"NotesAPI/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func TestGetAllReminders(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/notes/reminders", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)
	c.Set("user", jwt.MapClaims{
		"userId": float64(1),
	})
	err := GetAllRemindersController(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}

func TestCreateReminder(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/notes/reminders", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)
	c.Set("user", jwt.MapClaims{
		"userId": float64(1),
	})
	err := CreateReminderController(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}

func TestUpdateReminderReminder(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "/notes/reminders/:id", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)
	c.Set("user", jwt.MapClaims{
		"userId": float64(1),
	})
	err := UpdateReminderByIDController(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}

func TestGetReminderByID(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "/notes/reminders/:id", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)
	c.Set("user", jwt.MapClaims{
		"userId": float64(1),
	})
	err := GetReminderByIDController(c)
	// if err != nil {
	// 	t.Error(err.Error())
	// }

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}
func TestDeleteReminderByID(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/notes/reminders/:id", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)
	c.Set("user", jwt.MapClaims{
		"userId": float64(1),
	})
	err := DeleteReminderByIDController(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}
