package repository

import (
	"log"
	"os"

	"github.com/nepackage/nepackage/models"
	"github.com/nepackage/nepackage/utils"
)

func GetGitHubCredentials() ([]models.GithubCredential, error) {
	var githubCredentials []models.GithubCredential
	if err := models.DB.Find(&githubCredentials).Error; err != nil {
		return nil, err
	}
	return githubCredentials, nil
}

func GetGitHubCredentialById(ghCredentialId uint) (*models.GithubCredential, error) {
	var ghCredential *models.GithubCredential
	if err := models.DB.Where("id = ?", ghCredentialId).First(&ghCredential).Error; err != nil {
		return nil, err
	}
	return ghCredential, nil
}

func CreateGitHubCredential(ghCredential *models.GithubCredentialCreate) (*models.GithubCredential, error) {
	if err := models.DB.Table("github_credentials").Create(&ghCredential).Error; err != nil {
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

func UpdateGitHubCredential(ghCredentialId uint, ghCredentialInput models.GithubCredentialUpdate) (*models.GithubCredential, error) {
	var ghCredential models.GithubCredential
	if ghCredentialInput.Password != "" {
		newEcryptedPassword, err := utils.EncryptString(ghCredentialInput.Password, os.Getenv("JWTKEY"))
		if err != nil {
			log.Println("error encrypting new password ", err.Error())
			return nil, err
		}
		ghCredentialInput.Password = newEcryptedPassword
	}
	if err := models.DB.Table("github_credentials").Where("id = ?", ghCredentialId).Model(&ghCredential).Updates(ghCredentialInput).Error; err != nil {
		log.Println("error updating github credential", err.Error())
		return nil, err
	}
	ghCredentialUpdated := models.GithubCredential{
		ID:        ghCredential.ID,
		Name:      ghCredential.Name,
		Username:  ghCredential.Username,
		State:     ghCredential.State,
		CreatedAt: ghCredential.CreatedAt,
		UpdatedAt: ghCredential.UpdatedAt,
	}
	log.Println("github credential updated")
	return &ghCredentialUpdated, nil
}
