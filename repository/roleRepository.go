package repository

import (
	"github.com/nepackage/nepackage/database"
	"github.com/nepackage/nepackage/models"
)

func GetRoles() ([]models.Role, error) {
	var roles []models.Role
	if err := database.DB.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func GetRoleById(roleId uint) (*models.Role, error) {
	var role *models.Role
	if err := database.DB.Where("id = ?", roleId).First(&role).Error; err != nil {
		return nil, err
	}

	return role, nil
}

func CreateRole(role *models.Role) (bool, error) {
	if err := database.DB.Create(&role).Error; err != nil {
		return false, err
	}

	return true, nil
}

func UpdateRole(roleId uint, roleInput models.RoleUpdate) (*models.RoleUpdate, error) {
	var role models.Role

	if err := database.DB.Where("id = ?", roleId).Model(&role).Updates(roleInput).Error; err != nil {
		return nil, err
	}

	return &roleInput, nil
}

func DeleteRole(roleId uint) (bool, error) {
	var role models.Role
	if err := database.DB.Where("id = ?", roleId).First(&role).Error; err != nil {
		return false, err
	}
	if err := database.DB.Delete(&role).Error; err != nil {
		return false, err
	}

	return true, nil
}
