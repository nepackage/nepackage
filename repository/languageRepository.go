package repository

import (
	"github.com/nepackage/nepackage/models"
)

func GetLanguages() ([]models.Language, error) {
	var languages []models.Language
	if err := models.DB.Find(&languages).Error; err != nil {
		return nil, err
	}
	return languages, nil
}

func GetLanguageById(languageId uint) (*models.Language, error) {
	var language *models.Language
	if err := models.DB.Where("id = ?", languageId).First(&language).Error; err != nil {
		return nil, err
	}

	return language, nil
}

func CreateLanguage(language *models.Language) (bool, error) {
	if err := models.DB.Create(&language).Error; err != nil {
		return false, err
	}

	return true, nil
}

func UpdateLanguage(languageId uint, languageInput models.LanguageUpdate) (*models.LanguageUpdate, error) {
	var language models.Language

	if err := models.DB.Where("id = ?", languageId).Model(&language).Updates(languageInput).Error; err != nil {
		return nil, err
	}

	return &languageInput, nil
}

func DeleteLanguage(languageId uint) (bool, error) {
	var language models.Language
	if err := models.DB.Where("id = ?", languageId).First(&language).Error; err != nil {
		return false, err
	}
	if err := models.DB.Delete(&language).Error; err != nil {
		return false, err
	}

	return true, nil
}
