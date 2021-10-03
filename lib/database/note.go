package database

import (
	"NotesAPI/config"
	"NotesAPI/model"
	"fmt"
	"time"
)

func GetNotes() []model.Note {
	var Notes []model.Note
	config.DB.Debug().Preload("Labels").Where("deleted_at is null").Find(&Notes)
	return Notes
}

func GetNotesByID(noteId int) (interface{}, error) {
	var note model.Note
	if e := config.DB.Where("deleted_at is null").Find(&note, noteId).Error; e != nil {
		return nil, e
	}
	return note, nil
}

func CreateNote(note model.Note) model.Note {
	config.DB.Create(&note)
	for _, label := range note.Labels {
		err := config.DB.Debug().Raw("insert into notelabels values(?)", label.ID).Error
		if err != nil {
			fmt.Println(err)
		}
	}
	return note
}

func DeleteNoteByID(id string) {
	var note model.Note
	deletNote := time.Now()
	note.DeletedAt = &deletNote
	config.DB.Where("id = ?", id).Updates(&note)
}

func UpdateNoteByID(id string, note model.Note) {
	config.DB.Where("id = ?", id).Updates(&note)
}
