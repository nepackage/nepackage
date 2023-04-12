package repository

import (
	"github.com/nepackage/nepackage/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User
	if err := models.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserById(userId uint) (*models.User, error) {
	var user *models.User
	if err := models.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func CreateUser(user *models.UserCreate) (bool, error) {
	if err := models.DB.Table("users").Create(&user).Error; err != nil {
		return false, err
	}

	return true, nil
}

func UpdateUser(userId uint, userInput models.UserUpdate) (*models.UserUpdate, error) {
	var user models.User

	if err := models.DB.Where("id = ?", userId).Model(&user).Updates(userInput).Error; err != nil {
		return nil, err
	}

	return &userInput, nil
}

func DeleteUser(userId uint) (bool, error) {
	var user models.User
	if err := models.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return false, err
	}
	if err := models.DB.Delete(&user).Error; err != nil {
		return false, err
	}

	return true, nil
}
