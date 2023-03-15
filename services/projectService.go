package services

import (
	"github.com/otisnado/nepackage/models"
)

func GetProjects() ([]models.Project, error) {
	var projects []models.Project
	if err := models.DB.Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func GetProjectById(projectId uint) (*models.Project, error) {
	var project *models.Project
	if err := models.DB.Where("id = ?", projectId).First(&project).Error; err != nil {
		return nil, err
	}

	return project, nil
}

func CreateProject(project *models.Project) (bool, error) {
	if err := models.DB.Create(&project).Error; err != nil {
		return false, err
	}

	return true, nil
}

func UpdateProject(projectId uint, projectInput models.ProjectUpdate) (*models.ProjectUpdate, error) {
	var project models.Project

	if err := models.DB.Where("id = ?", projectId).Model(&project).Updates(projectInput).Error; err != nil {
		return nil, err
	}

	return &projectInput, nil
}

func DeleteProject(projectId uint) (bool, error) {
	var project models.Project
	if err := models.DB.Where("id = ?", projectId).First(&project).Error; err != nil {
		return false, err
	}
	if err := models.DB.Delete(&project).Error; err != nil {
		return false, err
	}

	return true, nil
}
