package database

import (
	"NotesAPI/config"
	"NotesAPI/middleware"
	"NotesAPI/model"
)

func GetUsers() []model.User {
	var users []model.User
	config.DB.Preload("Notes").Find(&users)
	return users
}

func GetUserByID(id string) model.User {
	var user model.User
	config.DB.Where("id = ?", id).Preload("Notes").Find(&user)
	return user
}

func CreateUser(user model.User) model.User {
	config.DB.Create(&user)
	return user
}

func DeleteUserByID(id string) {
	var user model.User
	config.DB.Where("id = ?", id).Delete(&user)
}

func UpdateUserByID(id string, user model.User) {
	config.DB.Where("id = ?", id).Updates(&user)
}

func IsValid(nm string, em string) bool {
	var user model.User
	if err := config.DB.Where("email = ?", em).Find(&user).Error; err != nil {
		return false
	}

	return nm == user.Name
}

func GetDetailUser(userId int)(interface{}, error){
	var user model.User

	if e:=config.DB.Find(&user,userId).Error; e!=nil{
		return nil,e
	}
	return user,nil
}

func LoginUser(user *model.User)(interface{},error){
	var err error
	if err = config.DB.Where("name = ? AND email = ?", user.Name, user.Email).First(user).Error; err!=nil{
		return nil,err
	}

	user.Token, err= middleware.CreateToken(int(user.ID))
	if err != nil{
		return nil,err
	}

	if err:=config.DB.Save(user).Error; err!=nil{
		return nil,err
	}

	return user,nil
}
