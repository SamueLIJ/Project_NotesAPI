package model

import "time"

type Label struct {
	// gorm.Model
	ID         uint       `gorm:"primarykey" json:"id" form:"id"`
	Notes      []Note     `gorm:"many2many:notelabels" json:"labels,omitempty"`
	LabelTitle string     `json:"labeltitle" form:"labeltitle"`
	CreatedAt  time.Time  `json:"createdAt"`
	DeletedAt  *time.Time `json:"deletedAt"`
}
