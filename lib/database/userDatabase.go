package database

import (
	"myproperty-api/config"
	"myproperty-api/middlewares"
	model "myproperty-api/models"
)

func GetUsers() []model.User {
	var users []model.User
	config.DB.Preload("Wishlist").Find(&users)
	return users
}

func GetUserByID(id string) model.User {
	var user model.User
	config.DB.Where("id = ?", id).Preload("Wishlist").Find(&user)
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

func GetDetailUsers(userId int) (interface{}, error) {
	var user model.User

	if e := config.DB.Find(&user, userId).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func LoginUsers(user *model.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
