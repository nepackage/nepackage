package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nepackage/nepackage/models"
	"github.com/nepackage/nepackage/repository"
)

// FindLanguages		godoc
// @Summary			FindLanguages
// @Schemes
// @Description		Bulk all nepackage's languages
// @Tags			Languages
// @Produce			json
// @Param			Authorization		header	string	true	"JWT without bearer"
// @Success			200		{object}	models.SuccessFindLanguages
// @Failure			401,500	{object}	models.ErrorMessage
// @Router			/languages	[get]
func FindLanguages(c *gin.Context) {
	languages, err := repository.GetLanguages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": languages})
}

// FindLanguage		godoc
// @Summary		FindLanguage
// @Schemes
// @Description	Find a language with given id in path
// @Tags		Languages
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		id			path		int		true	"Language ID"
// @Success		200			{object}	models.SuccessFindLanguage
// @Failure		400,401,404	{object}	models.ErrorMessage
// @Router		/languages/{id}			[get]
func FindLanguage(c *gin.Context) {
	language_id, _ := strconv.Atoi(c.Param("id"))

	language, err := repository.GetLanguageById(uint(language_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": language})
}

// CreateLanguage	godoc
// @Summary		CreateLanguage
// @Schemes
// @Description	Create language with models.Language model
// @Tags		Languages
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		language	body		models.Language		true	"Language data"
// @Success		200			{object}	models.SuccessLanguageCreation
// @Failure		400,401,500	{object}	models.ErrorMessage
// @Router		/languages	[post]
func CreateLanguage(c *gin.Context) {
	var input models.Language
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := repository.CreateLanguage(&input)
	if !created {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": &input})
}

// UpdateLanguage	godoc
// @Summary		UpdateLanguage
// @Schemes
// @Description	Update language with models.Language model
// @Tags		Languages
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		language		body		models.Language			true	"Language data"
// @Success		200				{object}	models.SuccessLanguageUpdate
// @Failure		400,401,404,500	{object}	models.ErrorMessage
// @Router		/languages/{id}	[patch]
func UpdateLanguage(c *gin.Context) {
	language_id, _ := strconv.Atoi(c.Param("id"))

	var input models.LanguageUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := repository.GetLanguageById(uint(language_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	languageUpdated, err := repository.UpdateLanguage(uint(language_id), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": languageUpdated})
}

// DeleteLanguage	godoc
// @Summary		DeleteLanguage
// @Schemes
// @Description	Delete language with using id
// @Tags		Languages
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		id			path		int				true	"Language ID"
// @Success		200			{object}	models.SuccessLanguageDelete
// @Failure		401,404,500	{object}	models.ErrorMessage
// @Router		/languages/{id}	[delete]
func DeleteLanguage(c *gin.Context) {
	language_id, _ := strconv.Atoi(c.Param("id"))

	_, err := repository.GetLanguageById(uint(language_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	state, err := repository.DeleteLanguage(uint(language_id))
	if !state {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": state})
}
