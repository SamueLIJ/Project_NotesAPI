package controller

import (
	"NotesAPI/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func TestGetAllNotes(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/notes", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)
	c.Set("user", jwt.MapClaims{
		"userId": float64(1),
	})
	err := GetAllNotesController(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}

func TestCreateNote(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/notes", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)
	c.Set("user", jwt.MapClaims{
		"userId": float64(1),
	})
	err := CreateNoteController(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}

func TestUpdateNote(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "/notes/:id", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)
	c.Set("user", jwt.MapClaims{
		"userId": float64(1),
	})
	err := UpdateNoteByIDController(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}

func TestGetNoteByID(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "/notes/:id", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)
	c.Set("user", jwt.MapClaims{
		"userId": float64(1),
	})
	err := GetNoteByIDController(c)
	// if err != nil {
	// 	t.Error(err.Error())
	// }

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}
func TestDeleteNoteByID(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/notes/:id", nil)
	rec := httptest.NewRecorder()

	config.InitDB()
	app := echo.New()
	c := app.NewContext(req, rec)
	c.Set("user", jwt.MapClaims{
		"userId": float64(1),
	})
	err := DeleteNoteByIDController(c)
	if err != nil {
		t.Error(err.Error())
	}

	if rec.Result().StatusCode != 200 {
		t.Error(err.Error())
	}
}
