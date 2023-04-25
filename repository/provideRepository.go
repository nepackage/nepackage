package repository

import "github.com/nepackage/nepackage/models"

func GetAllProviders() ([]models.Provider, error) {
	var providers []models.Provider
	if err := models.DB.Find(&providers).Error; err != nil {
		return nil, err
	}
	return providers, nil
}
