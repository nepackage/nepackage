package repository

import (
	"errors"

	"github.com/nepackage/nepackage/database"
	"github.com/nepackage/nepackage/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserById(userId uint) (*models.User, error) {
	var user *models.User
	if err := database.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func CreateUser(user *models.UserCreate) (bool, error) {
	if err := database.DB.Table("users").Create(&user).Error; err != nil {
		return false, err
	}

	return true, nil
}

func UpdateUser(userId uint, userInput models.UserUpdate) (*models.UserUpdate, error) {
	var user models.User

	if err := database.DB.Where("id = ?", userId).Model(&user).Updates(userInput).Error; err != nil {
		return nil, err
	}

	return &userInput, nil
}

func DeleteUser(userId uint) (bool, error) {
	var user models.User
	if err := database.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return false, err
	}
	if err := database.DB.Delete(&user).Error; err != nil {
		return false, err
	}

	return true, nil
}

func CheckIfMailExists(mailInput string) error {
	var user models.User
	result := database.DB.Where("mail = ?", mailInput).Find(&user)
	if result.RowsAffected > 0 {
		return errors.New("mail exists")
	}
	return nil
}

func CheckIfUsernameExists(usernameInput string) error {
	var user models.User
	result := database.DB.Where("username = ?", usernameInput).Find(&user)
	if result.RowsAffected > 0 {
		return errors.New("username exists")
	}
	return nil
}
