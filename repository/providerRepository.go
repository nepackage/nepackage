package repository

import "github.com/nepackage/nepackage/models"

func GetAllProviders() ([]models.Provider, error) {
	var providers []models.Provider
	if err := models.DB.Find(&providers).Error; err != nil {
		return nil, err
	}
	return providers, nil
}

func GetProviderById(providerId uint) (*models.Provider, error) {
	var provider *models.Provider
	if err := models.DB.Where("id = ?", providerId).First(&provider).Error; err != nil {
		return nil, err
	}
	return provider, nil
}

func CreateProvider(provider *models.Provider) (*models.Provider, error) {
	if err := models.DB.Create(&provider).Error; err != nil {
		return nil, err
	}
	return provider, nil
}

func UpdateProvider(providerId uint, newData *models.Provider) (*models.Provider, error) {
	var provider models.Provider
	if err := models.DB.Where("id = ?", providerId).Model(provider).Updates(newData).Error; err != nil {
		return nil, err
	}
	if err := models.DB.Where("id = ?", providerId).First(&provider).Error; err != nil {
		return nil, err
	}
	return &provider, nil
}

func DeleteProvider(providerId uint) error {
	var provider models.Provider
	if err := models.DB.Where("id = ?", providerId).First(&provider).Error; err != nil {
		return err
	}
	if err := models.DB.Delete(&provider).Error; err != nil {
		return err
	}
	return nil
}
