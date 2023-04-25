package services

import (
	"github.com/nepackage/nepackage/models"
	"github.com/nepackage/nepackage/repository"
)

func GetProviders() ([]models.Provider, error) {
	providers, err := repository.GetAllProviders()
	if err != nil {
		return nil, err
	}
	return providers, nil
}

func GetProviderById(providerId uint) (*models.Provider, error) {
	provider, err := repository.GetProviderById(providerId)
	if err != nil {
		return nil, err
	}
	return provider, nil
}

func CreateProvider(input *models.Provider) (*models.Provider, error) {
	newProvider, err := repository.CreateProvider(input)
	if err != nil {
		return nil, err
	}
	return newProvider, nil
}

func UpdateProvider(providerId uint, newData *models.Provider) (*models.Provider, error) {
	updatedProvider, err := repository.UpdateProvider(providerId, newData)
	if err != nil {
		return nil, err
	}
	return updatedProvider, nil
}

func DeleteProvider(providerId uint) error {
	err := repository.DeleteProvider(providerId)
	if err != nil {
		return err
	}
	return nil
}
