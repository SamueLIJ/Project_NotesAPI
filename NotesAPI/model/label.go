package model

import "time"

type Label struct {
	// gorm.Model
	ID         uint   `gorm:"primarykey" form:"id"`
	Notes      []Note `gorm:"many2many:notelabels" json:"labels,omitempty"`
	LabelTitle string `json:"labeltitle" form:"labeltitle"`
	CreatedAt  time.Time
}
