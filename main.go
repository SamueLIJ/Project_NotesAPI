package main

import (
	"NotesAPI/config"
	"NotesAPI/middleware"
	"NotesAPI/route"

	"github.com/labstack/echo"
)

func main() {
	config.InitDB()
	config.InitLog()
	config.InitMigration()

	e := echo.New()
	e.Use(middleware.Log)
	route.NewUser(e)
	route.NewNotes(e)
	route.NewLabel(e)
	route.NewReminders(e)
	route.NewPictures(e)
	e.Start(":8080")
}
