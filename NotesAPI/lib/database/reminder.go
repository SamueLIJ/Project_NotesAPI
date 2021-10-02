package database

import (
	"NotesAPI/config"
	"NotesAPI/model"
)

func GetReminders() []model.Reminder {
	var reminders []model.Reminder
	config.DB.Find(&reminders)
	return reminders
}

func GetReminderByID(id string) model.Reminder {
	var reminder model.Reminder
	config.DB.Where("id = ?", id).Find(&reminder)
	return reminder
}

func CreateReminder(reminder model.Reminder) model.Reminder {
	config.DB.Create(&reminder)
	return reminder
}

func DeleteReminderByID(id string) {
	var reminder model.Reminder
	config.DB.Where("id = ?", id).Delete(&reminder)
}

func UpdateReminderByID(id string, reminder model.Reminder) {
	config.DB.Where("id = ?", id).Updates(&reminder)
}
