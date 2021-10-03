package model

import "time"

type Reminder struct {
	// gorm.Model
	ID            uint       `gorm:"primarykey" form:"id"`
	Reminder_name string     `json:"reminder_name"`
	Reminder_time time.Time  `json:"reminder_time"`
	CreatedAt     time.Time  `json:"createdAt"`
	Note          *Note      `json:"note,omitempty"`
	NoteID        uint       `json:"noteid"`
	DeletedAt     *time.Time `json:"deletedAt"`
}
