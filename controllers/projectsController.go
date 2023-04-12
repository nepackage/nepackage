package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nepackage/nepackage/models"
	"github.com/nepackage/nepackage/repository"
	"github.com/nepackage/nepackage/services"
)

// FindProjects		godoc
// @Summary			FindProjects
// @Schemes
// @Description		Bulk all nepackage's projects
// @Tags			Projects
// @Produce			json
// @Param			Authorization		header	string	true	"JWT without bearer"
// @Success			200		{object}	models.SuccessFindProjects
// @Failure			401,500	{object}	models.ErrorMessage
// @Router			/projects	[get]
func FindProjects(c *gin.Context) {
	projects, err := repository.GetProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": projects})
}

// FindProject	godoc
// @Summary		FindProject
// @Schemes
// @Description	Find a project with given id in path
// @Tags		Projects
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		id			path		int		true	"Project ID"
// @Success		200			{object}	models.SuccessFindProject
// @Failure		400,401,404	{object}	models.ErrorMessage
// @Router		/projects/{id}			[get]
func FindProject(c *gin.Context) {
	project_id, _ := strconv.Atoi(c.Param("id"))

	project, err := repository.GetProjectById(uint(project_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": project})
}

// CreateProject	godoc
// @Summary		CreateProject
// @Schemes
// @Description	Create project with models.Project model
// @Tags		Projects
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		project		body		models.Project		true	"Project data"
// @Success		200			{object}	models.SuccessProjectCreation
// @Failure		400,401,500	{object}	models.ErrorMessage
// @Router		/projects	[post]
func CreateProject(c *gin.Context) {
	var input models.Project
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := repository.CreateProject(&input)
	if !created {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": &input})
}

// UpdateProject	godoc
// @Summary		UpdateProject
// @Schemes
// @Description	Update project with models.Project model
// @Tags		Projects
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		project			body		models.Project			true	"Project data"
// @Success		200				{object}	models.SuccessProjectUpdate
// @Failure		400,401,404,500	{object}	models.ErrorMessage
// @Router		/projects/{id}	[patch]
func UpdateProject(c *gin.Context) {
	project_id, _ := strconv.Atoi(c.Param("id"))

	var input models.ProjectUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := repository.GetProjectById(uint(project_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	projectUpdated, err := repository.UpdateProject(uint(project_id), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": projectUpdated})
}

// DeleteProject	godoc
// @Summary		DeleteProject
// @Schemes
// @Description	Delete project with using id
// @Tags		Projects
// @Produce		json
// @Param		Authorization		header	string	true	"JWT without bearer"
// @Param		id			path		int				true	"Project ID"
// @Success		200			{object}	models.SuccessProjectDelete
// @Failure		401,404,500	{object}	models.ErrorMessage
// @Router		/projects/{id}	[delete]
func DeleteProject(c *gin.Context) {
	project_id, _ := strconv.Atoi(c.Param("id"))

	_, err := repository.GetProjectById(uint(project_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	state, err := repository.DeleteProject(uint(project_id))
	if !state {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": state})
}

func CreateSpringProject(c *gin.Context) {
	var springProjectInput models.SpringProject
	if err := c.ShouldBindJSON(&springProjectInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, projectPath, err := services.SpringProjectGenerator(springProjectInput)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err})
		return
	}
	err = services.InitLocalRepository(projectPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = services.FirstCommitLocalRepository(projectPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	repositoryUrl, err := services.CreateGitHubRemoteRepository(springProjectInput.Name, springProjectInput.Visibility)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = services.PushLocalRepositoryToGitHub(projectPath, repositoryUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": springProjectInput})
}
