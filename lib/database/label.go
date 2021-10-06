package database

import (
	"NotesAPI/config"
	"NotesAPI/model"
	"fmt"
	"time"
)

func GetLabels() []model.Label {
	var Labels []model.Label

	config.DB.Where("deleted_at is null").Find(&Labels)
	return Labels
}

func GetLabelsByID(labelId int) (interface{}, error) {
	var label model.Label

	if e := config.DB.Where("deleted_at is null").Find(&label, labelId).Error; e != nil {
		return nil, e
	}
	return label, nil
}

func CreateLabel(label model.Label) model.Label {
	err:= config.DB.Create(&label).Error
	fmt.Println(err)
	for _, note := range label.Notes {
		config.DB.Raw("insert into notelabels values(?,?)", note.ID, label.ID)
	}
	return label
}

func DeleteLabelByID(id string) {
	var label model.Label
	deletLabel := time.Now()
	label.DeletedAt = &deletLabel
	config.DB.Where("id = ?", id).Updates(&label)
}

func UpdateLabelByID(id string, label model.Label) {
	config.DB.Where("id = ?", id).Updates(&label)
}
