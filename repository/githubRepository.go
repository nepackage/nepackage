package repository

import (
	"log"
	"os"

	"github.com/nepackage/nepackage/database"
	"github.com/nepackage/nepackage/models"
	"github.com/nepackage/nepackage/utils"
	"gorm.io/gorm/clause"
)

func GetGitHubCredentials() ([]models.GithubCredential, error) {
	var githubCredentials []models.GithubCredential
	if err := database.DB.Find(&githubCredentials).Error; err != nil {
		return nil, err
	}
	return githubCredentials, nil
}

func GetGitHubCredentialById(ghCredentialId uint) (*models.GithubCredential, error) {
	var ghCredential *models.GithubCredential
	if err := database.DB.Where("id = ?", ghCredentialId).First(&ghCredential).Error; err != nil {
		return nil, err
	}
	return ghCredential, nil
}

func CreateGitHubCredential(ghCredential *models.GithubCredentialCreate) (*models.GithubCredential, error) {
	if err := database.DB.Table("github_credentials").Create(&ghCredential).Error; err != nil {
		return nil, err
	}

	ghCredentialCreated := models.GithubCredential{
		ID:        ghCredential.ID,
		Name:      ghCredential.Name,
		Username:  ghCredential.Username,
		Password:  ghCredential.Password,
		State:     ghCredential.State,
		CreatedAt: ghCredential.CreatedAt,
		UpdatedAt: ghCredential.UpdatedAt,
	}

	return &ghCredentialCreated, nil
}

func UpdateGitHubCredential(ghCredentialId uint, ghCredentialInput models.GithubCredentialUpdate) (*models.GithubCredentialUpdate, error) {
	var ghCredential models.GithubCredential
	if ghCredentialInput.Password != "" {
		newEcryptedPassword, err := utils.EncryptString(ghCredentialInput.Password, os.Getenv("JWTKEY"))
		if err != nil {
			log.Println("error encrypting new password ", err.Error())
			return nil, err
		}
		ghCredentialInput.Password = newEcryptedPassword
	}
	if err := database.DB.Table("github_credentials").Clauses(clause.Returning{}).Where("id = ?", ghCredentialId).Model(&ghCredential).Updates(ghCredentialInput).Error; err != nil {
		log.Println("error updating github credential", err.Error())
		return nil, err
	}

	log.Println("github credential updated with Id: ", ghCredentialId, " was updated")
	return &ghCredentialInput, nil
}

func DeleteGitHubCredential(ghCredentialId uint) (bool, error) {
	var ghCredential models.GithubCredential
	if err := database.DB.Where("id = ?", ghCredentialId).First(&ghCredential).Error; err != nil {
		return false, err
	}
	if err := database.DB.Delete(&ghCredential).Error; err != nil {
		return false, err
	}

	return true, nil
}
