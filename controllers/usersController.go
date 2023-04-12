package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nepackage/nepackage/models"
	"github.com/nepackage/nepackage/repository"
	"github.com/nepackage/nepackage/utils"
)

// FindUsers		godoc
// @Summary			FindUsers
// @Schemes
// @Description		Bulk all nepackage's users
// @Tags			Users
// @Produce			json
// @Param			Authorization		header	string	true	"JWT without bearer"
// @Success			200		{object}	models.SuccessFindUsers
// @Failure			401,500	{object}	models.ErrorMessage
// @Router			/users	[get]
func FindUsers(c *gin.Context) {
	users, err := repository.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// FindUser		godoc
// @Summary		FindUser
// @Schemes
// @Description	Find a user with given id in path
// @Tags		Users
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		id			path		int		true	"User ID"
// @Success		200			{object}	models.SuccessFindUser
// @Failure		400,401,404	{object}	models.ErrorMessage
// @Router		/users/{id}			[get]
func FindUser(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Param("id"))

	user, err := repository.GetUserById(uint(user_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// CreateUser	godoc
// @Summary		CreateUser
// @Schemes
// @Description	Create user with models.User model
// @Tags		Users
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		user		body		models.User		true	"User data"
// @Success		200			{object}	models.SuccessUserCreation
// @Failure		400,401,500	{object}	models.ErrorMessage
// @Router		/users	[post]
func CreateUser(c *gin.Context) {
	var input models.UserCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.CheckIfMailExists(input.Mail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.CheckIfUsernameExists(input.Username); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := input.HashPassword(input.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	created, err := repository.CreateUser(&input)
	if !created {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": &input})
}

// UpdateUser	godoc
// @Summary		UpdateUser
// @Schemes
// @Description	Update user with models.User model
// @Tags		Users
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		user			body		models.User			true	"User data"
// @Success		200				{object}	models.SuccessUserUpdate
// @Failure		400,401,404,500	{object}	models.ErrorMessage
// @Router		/users/{id}	[patch]
func UpdateUser(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Param("id"))

	var input models.UserUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := repository.GetUserById(uint(user_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	userUpdated, err := repository.UpdateUser(uint(user_id), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": userUpdated})
}

// DeleteUser	godoc
// @Summary		DeleteUser
// @Schemes
// @Description	Delete user with using id
// @Tags		Users
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		id			path		int				true	"User ID"
// @Success		200			{object}	models.SuccessUserDelete
// @Failure		401,404,500	{object}	models.ErrorMessage
// @Router		/users/{id}	[delete]
func DeleteUser(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Param("id"))

	_, err := repository.GetUserById(uint(user_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	state, err := repository.DeleteUser(uint(user_id))
	if !state {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": state})
}
