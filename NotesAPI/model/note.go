package model

import (
	"time"
)

type Note struct {
	// gorm.Model
	ID           uint        `gorm:"primarykey" form:"id"`
	UserID       uint        `json:"userid"`
	User         *User       `json:"user,omitempty"`
	Title        string      `json:"title" form:"title"`
	Content      string      `json:"content" form:"content"`
	Status       *string     `json:"status" form:"status"`
	Archive_date *time.Time  `json:"archive_date" form:"archive_date"`
	Audiofile    *string     `json:"audiofile"`
	Labels       []Label     `gorm:"many2many:notelabels" json:"labels,omitempty"`
	Reminders    []*Reminder `json:"reminder"`
	Pictures     []*Picture  `json:"picture"`
	CreatedAt    time.Time   `json:"createdAt"`
	DeletedAt    *time.Time  `json:"deletedAt"`
}
