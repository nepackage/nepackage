package utils

import (
	"errors"

	"github.com/nepackage/nepackage/models"
)

func CheckIfMailExists(mailInput string) error {
	var user models.User
	result := models.DB.Where("mail = ?", mailInput).Find(&user)
	if result.RowsAffected > 0 {
		return errors.New("mail exists")
	}
	return nil
}

func CheckIfUsernameExists(usernameInput string) error {
	var user models.User
	result := models.DB.Where("username = ?", usernameInput).Find(&user)
	if result.RowsAffected > 0 {
		return errors.New("username exists")
	}
	return nil
}
