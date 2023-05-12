package repository

import (
	"github.com/nepackage/nepackage/database"
	"github.com/nepackage/nepackage/models"
)

func GetGroups() ([]models.Group, error) {
	var groups []models.Group
	if err := database.DB.Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

func GetGroupById(groupId uint) (*models.Group, error) {
	var group *models.Group
	if err := database.DB.Where("id = ?", groupId).First(&group).Error; err != nil {
		return nil, err
	}

	return group, nil
}

func CreateGroup(group *models.Group) (bool, error) {
	if err := database.DB.Create(&group).Error; err != nil {
		return false, err
	}

	return true, nil
}

func UpdateGroup(groupId uint, groupInput models.GroupUpdate) (*models.GroupUpdate, error) {
	var group models.Group

	if err := database.DB.Where("id = ?", groupId).Model(&group).Updates(groupInput).Error; err != nil {
		return nil, err
	}

	return &groupInput, nil
}

func DeleteGroup(groupId uint) (bool, error) {
	var group models.Group
	if err := database.DB.Where("id = ?", groupId).First(&group).Error; err != nil {
		return false, err
	}
	if err := database.DB.Delete(&group).Error; err != nil {
		return false, err
	}

	return true, nil
}
