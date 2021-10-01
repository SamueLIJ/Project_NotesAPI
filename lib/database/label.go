package database

import (
	"NotesAPI/config"
	"NotesAPI/model"
)

func GetLabels() []model.Label {
	var Labels []model.Label
	config.DB.Find(&Labels)
	return Labels
}

func GetLabelsByID(id string) model.Label {
	var label model.Label
	config.DB.Where("id = ?", id).Find(&label)
	return label
}

func CreateLabel(label model.Label) model.Label {
	config.DB.Create(&label)
	for _, note := range label.Notes {
		config.DB.Raw("insert into notelabels values(?,?)", note.ID)
	}
	return label
}

func DeleteLabelByID(id string) {
	var label model.Label
	config.DB.Where("id = ?", id).Delete(&label)
}

func UpdateLabelByID(id string, label model.Label) {
	config.DB.Where("id = ?", id).Updates(&label)
}
