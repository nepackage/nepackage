package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nepackage/nepackage/models"
	"github.com/nepackage/nepackage/repository"
	"github.com/nepackage/nepackage/services"
)

func FindGitHubCredentials(c *gin.Context) {
	ghCredentials, err := services.FindAllGhCredentials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": ghCredentials})
}

func FindGitHubCredentialById(c *gin.Context) {
	ghCredentialId, _ := strconv.Atoi(c.Param("id"))
	ghCredential, err := services.FindAllGhCredentialById(uint(ghCredentialId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ghCredential})
}

func CreateGitHubCredential(c *gin.Context) {
	var input models.GithubCredentialCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ghCredentialCreated, err := services.CreateGitHubCredential(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": ghCredentialCreated})
}

func UpdateGitHubCredential(c *gin.Context) {
	var input models.GithubCredentialUpdate
	ghCredentialId, _ := strconv.Atoi(c.Param("id"))

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ghCredentialUpdated, err := services.UpdateGitHubCredential(uint(ghCredentialId), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": &ghCredentialUpdated})
}

func DeleteGitHubCredential(c *gin.Context) {
	ghCredentialId, _ := strconv.Atoi(c.Param("id"))

	_, err := repository.GetGitHubCredentialById(uint(ghCredentialId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	state, err := services.DeleteGitHubCredential(uint(ghCredentialId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": state})
}
