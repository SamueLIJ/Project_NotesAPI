package model

import (
	"time"
)

type Note struct {
	// gorm.Model
	ID           uint        `gorm:"primarykey" form:"id"`
	UserID       uint        `json:"userid" form:"userid"`
	User         *User       `json:"user,omitempty" form:"user"`
	Title        string      `json:"title" form:"title"`
	Content      string      `json:"content" form:"content"`
	Status       *string     `json:"status" form:"status"`
	Archive_date *time.Time  `json:"archive_date" form:"archive_date"`
	Audiofile    *string     `json:"audiofile" form:"audio_file"`
	Labels       []Label     `gorm:"many2many:notelabels" json:"labels,omitempty"`
	Reminders    []*Reminder `json:"reminder" form:"reminder"`
	Pictures     []*Picture  `json:"picture" form:"picture"`
	CreatedAt    time.Time   `json:"createdAt" form:"createdAt"`
	DeletedAt    *time.Time  `json:"deletedAt" form:"deletedAt"`
}
