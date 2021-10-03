package model

import "time"

type Picture struct {
	// gorm.Model
	ID           uint       `gorm:"primarykey" form:"id"`
	Picture_name string     `json:"picture" form:"picture"`
	CreatedAt    time.Time  `json:"createdAt" form:"createdAt"`
	Note         *Note      `json:"note,omitempty" form:"note,omitempty"`
	NoteID       uint       `json:"noteid" form:"noteid"`
	DeletedAt    *time.Time `json:"deletedAt"`
}
