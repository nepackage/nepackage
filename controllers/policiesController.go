package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nepackage/nepackage/models"
	"github.com/nepackage/nepackage/repository"
)

// FindPolicies		godoc
// @Summary			FindPolicies
// @Schemes
// @Description		Bulk all nepackage's policies
// @Tags			Policies
// @Produce			json
// @Param			Authorization		header	string	true	"JWT without bearer"
// @Success			200		{object}	models.SuccessFindPolicies
// @Failure			401,500	{object}	models.ErrorMessage
// @Router			/policies	[get]
func FindPolicies(c *gin.Context) {
	projects, err := repository.GetPolicies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": projects})
}

// FindPolicy	godoc
// @Summary		FindPolicy
// @Schemes
// @Description	Find a policy with given id in path
// @Tags		Policies
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		id			path		int		true	"Policy ID"
// @Success		200			{object}	models.SuccessFindPolicies
// @Failure		400,401,404	{object}	models.ErrorMessage
// @Router		/policies/{id}			[get]
func FindPolicy(c *gin.Context) {
	project_id, _ := strconv.Atoi(c.Param("id"))

	project, err := repository.GetPolicyById(uint(project_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": project})
}

// CreatePolicy	godoc
// @Summary		CreatePolicy
// @Schemes
// @Description	Create policy with models.Policy model
// @Tags		Policies
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		policy		body		models.Policy		true	"Policy data"
// @Success		200			{object}	models.SuccessPolicyCreation
// @Failure		400,401,500	{object}	models.ErrorMessage
// @Router		/policies	[post]
func CreatePolicy(c *gin.Context) {
	var input models.Policy
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := repository.CreatePolicy(&input)
	if !created {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": &input})
}

// UpdatePolicy	godoc
// @Summary		UpdatePolicy
// @Schemes
// @Description	Update policy with models.Policy model
// @Tags		Policies
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		policy			body		models.Policy			true	"Policy data"
// @Success		200				{object}	models.SuccessPolicyUpdate
// @Failure		400,401,404,500	{object}	models.ErrorMessage
// @Router		/policies/{id}	[patch]
func UpdatePolicy(c *gin.Context) {
	project_id, _ := strconv.Atoi(c.Param("id"))

	var input models.PolicyUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := repository.GetPolicyById(uint(project_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	projectUpdated, err := repository.UpdatePolicy(uint(project_id), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": projectUpdated})
}

// DeletePolicy	godoc
// @Summary		DeletePolicy
// @Schemes
// @Description	Delete policy with using id
// @Tags		Policies
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		id			path		int				true	"Policy ID"
// @Success		200			{object}	models.SuccessPolicyDelete
// @Failure		401,404,500	{object}	models.ErrorMessage
// @Router		/policies/{id}	[delete]
func DeletePolicy(c *gin.Context) {
	project_id, _ := strconv.Atoi(c.Param("id"))

	_, err := repository.GetPolicyById(uint(project_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	state, err := repository.DeletePolicy(uint(project_id))
	if !state {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": state})
}
