package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nepackage/nepackage/models"
	"github.com/nepackage/nepackage/repository"
)

// FindGroups		godoc
// @Summary			FindGroups
// @Schemes
// @Description		Bulk all nepackage's groups
// @Tags			Groups
// @Produce			json
// @Param			Authorization		header	string	true	"JWT without bearer"
// @Success			200		{object}	models.SuccessFindGroups
// @Failure			401,500	{object}	models.ErrorMessage
// @Router			/groups	[get]
func FindGroups(c *gin.Context) {
	groups, err := repository.GetGroups()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": groups})
}

// FindGroup	godoc
// @Summary		FindGroup
// @Schemes
// @Description	Find a group with given id in path
// @Tags		Groups
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		id			path		int		true	"Group ID"
// @Success		200			{object}	models.SuccessFindGroup
// @Failure		400,401,404	{object}	models.ErrorMessage
// @Router		/groups/{id}			[get]
func FindGroup(c *gin.Context) {
	group_id, _ := strconv.Atoi(c.Param("id"))

	group, err := repository.GetGroupById(uint(group_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": group})
}

// CreateGroups	godoc
// @Summary			CreateGroup
// @Schemes
// @Description		Create group with models.Group model
// @Tags			Groups
// @Produce			json
// @Param			Authorization		header	string	true	"JWT without bearer"
// @Param			group		body		models.Group	true	"Group data"
// @Success			200			{object}	models.SuccessGroupCreation
// @Failure			400,401,500	{object}	models.ErrorMessage
// @Router			/groups	[post]
func CreateGroup(c *gin.Context) {
	var input models.Group
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := repository.CreateGroup(&input)
	if !created {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": &input})
}

// UpdateGroup	godoc
// @Summary			UpdateGroup
// @Schemes
// @Description		Update group with models.Group model
// @Tags			Groups
// @Produce			json
// @Param			Authorization		header	string	true	"JWT without bearer"
// @Param			group			body		models.Group			true	"Group data"
// @Success			200				{object}	models.SuccessGroupUpdate
// @Failure			400,401,404,500	{object}	models.ErrorMessage
// @Router			/groups	[patch]
func UpdateGroup(c *gin.Context) {
	group_id, _ := strconv.Atoi(c.Param("id"))

	var input models.GroupUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := repository.GetGroupById(uint(group_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	groupUpdated, err := repository.UpdateGroup(uint(group_id), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": groupUpdated})
}

// DeleteGroup	godoc
// @Summary			DeleteGroup
// @Schemes
// @Description		Delete group using id
// @Tags			Groups
// @Produce			json
// @Param			Authorization		header	string	true	"JWT without bearer"
// @Param			id			path		int				true	"Group ID"
// @Success			200			{object}	models.SuccessGroupDelete
// @Failure			401,404,500	{object}	models.ErrorMessage
// @Router			/groups/{id}	[delete]
func DeleteGroup(c *gin.Context) {
	group_id, _ := strconv.Atoi(c.Param("id"))

	_, err := repository.GetGroupById(uint(group_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	state, err := repository.DeleteGroup(uint(group_id))
	if !state {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": state})
}
