package config

import (
	"NotesAPI/model"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DBLog *mongo.Client

func InitLog() {
	var err error

	DBLog, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/NoteAPILogs"))

	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = DBLog.Connect(ctx)

	if err != nil {
		panic(err)
	}

	DBLog.ListDatabaseNames(ctx, bson.M{})

}

func InitDB() {
	var err error
	//                              username:password@tcp(host:port)/nama_database
	DB, err = gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/notesapi?parseTime=True"))
	if err != nil {
		panic(err)
	} else {
		fmt.Println("DB connected")
	}
}

func InitMigration() {
	DB.AutoMigrate(
		&model.User{},
		&model.Note{},
		&model.Label{},
		&model.Reminder{},
		&model.Picture{},
	)

}
