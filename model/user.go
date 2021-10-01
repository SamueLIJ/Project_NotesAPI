package model

type User struct {
	// gorm.Model
	ID       uint   `gorm:"primarykey"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Notes  []*Note
	Token string `json:"token"`
}
