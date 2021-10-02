package middleware

import (
	"NotesAPI/config"
	"fmt"
	"time"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

func Log(next echo.HandlerFunc) echo.HandlerFunc {
	collection := config.DBLog.Database("NoteAPILogs").Collection("logs")

	return func(c echo.Context) error {

		log := bson.M{
			"time":   time.Now(),
			"method": c.Request().Method,
			"path":   c.Path(),
		}

		collection.InsertOne(c.Request().Context(), log)

		fmt.Println(log)

		return next(c)
	}
}
