package database

import (
	"NotesAPI/config"
	"NotesAPI/model"
)

func GetPictures() []model.Picture {
	var pictures []model.Picture
	config.DB.Find(&pictures)
	return pictures
}

func GetPictureByID(id string) model.Picture {
	var picture model.Picture
	config.DB.Where("id = ?", id).Find(&picture)
	return picture
}

func CreatePicture(picture model.Picture) model.Picture {
	config.DB.Create(&picture)
	return picture
}

func DeletePictureByID(id string) {
	var picture model.Picture
	config.DB.Where("id = ?", id).Delete(&picture)
}

func UpdatePictureByID(id string, picture model.Picture) {
	config.DB.Where("id = ?", id).Updates(&picture)
}
