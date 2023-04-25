package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nepackage/nepackage/models"
	"github.com/nepackage/nepackage/services"
)

func FindProviders(c *gin.Context) {
	providers, err := services.GetProviders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": providers})
}

func FindProviderById(c *gin.Context) {
	providerId, _ := strconv.Atoi(c.Param("id"))
	provider, err := services.GetProviderById(uint(providerId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": provider})
}

func CreateProvider(c *gin.Context) {
	var input models.Provider
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newProvider, err := services.CreateProvider(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": newProvider})
}

func UpdateProvider(c *gin.Context) {
	providerId, _ := strconv.Atoi(c.Param("id"))
	var input models.Provider
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := services.GetProviderById(uint(providerId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	updatedProvider, err := services.UpdateProvider(uint(providerId), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": updatedProvider})
}

func DeleteProvider(c *gin.Context) {
	providerId, _ := strconv.Atoi(c.Param("id"))
	_, err := services.GetProviderById(uint(providerId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = services.DeleteProvider(uint(providerId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"status": "deleted"})
}
