package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nepackage/nepackage/controllers"
	"github.com/nepackage/nepackage/docs"
	"github.com/nepackage/nepackage/middlewares"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	api := router.Group("/api/v1")
	{

		/* JWT */

		// Generate token
		router.POST("/api/v1/token", controllers.GenerateToken)

		// Refresh JWT
		router.POST("/api/v1/token/refresh", controllers.RefreshToken)

		secure := api.Group("/").Use(middlewares.Auth())
		{

			// Find all projects
			secure.GET("/projects", controllers.FindProjects)

			// Find a specific project --> id is required
			secure.GET("/projects/:id", controllers.FindProject)

			// Create a project
			secure.POST("/projects", controllers.CreateProject)

			// Update data for a project --> id is required
			secure.PATCH("/projects/:id", controllers.UpdateProject)

			// Delete a project --> id is required
			secure.DELETE("/projects/:id", controllers.DeleteProject)

			// Create springProject
			secure.POST("/projects/springProject", controllers.CreateSpringProject)

			/* Languages routes */

			// Find all languages supported
			secure.GET("/languages", controllers.FindLanguages)

			// Find a specific language by its id
			secure.GET("languages/:id", controllers.FindLanguage)

			// Create a language
			secure.POST("/languages", controllers.CreateLanguage)

			// Update data for a language --> id is required
			secure.PATCH("/languages/:id", controllers.UpdateLanguage)

			// Delete language using id
			secure.DELETE("languages/:id", controllers.DeleteLanguage)

			/* Users routes */

			// Find all users registered
			secure.GET("/users", controllers.FindUsers)

			// Find a specific user by its id
			secure.GET("/users/:id", controllers.FindUser)

			// Create a user
			secure.POST("/users", controllers.CreateUser)

			// Update data for a user --> id is required
			secure.PATCH("/users/:id", controllers.UpdateUser)

			// Delete a user --> id is required
			secure.DELETE("/users/:id", controllers.DeleteUser)

			/* Groups routes */

			// Find all groups registered
			secure.GET("/groups", controllers.FindGroups)

			// Find a specific group by its id
			secure.GET("/groups/:id", controllers.FindGroup)

			// Create a group
			secure.POST("/groups", controllers.CreateGroup)

			// Update data for a group --> id is required
			secure.PATCH("/groups/:id", controllers.UpdateGroup)

			// Delete a group --> id is required
			secure.DELETE("/groups/:id", controllers.DeleteGroup)

			/* Roles routes */

			// Find all roles registered
			secure.GET("/roles", controllers.FindRoles)

			// Find a specific role by its id
			secure.GET("/roles/:id", controllers.FindRole)

			// Create a role
			secure.POST("/roles", controllers.CreateRole)

			// Update data for a role --> id is required
			secure.PATCH("/roles/:id", controllers.UpdateRole)

			// Delete a role --> id is required
			secure.DELETE("/roles/:id", controllers.DeleteRole)

			/* Policies routes */

			// Find all policies registered
			secure.GET("/policies", controllers.FindPolicies)

			// Find a specific policy by its id
			secure.GET("/policies/:id", controllers.FindPolicy)

			// Create a policy
			secure.POST("/policies", controllers.CreatePolicy)

			// Update data for a policy --> id is required
			secure.PATCH("/policies/:id", controllers.UpdatePolicy)

			// Delete a policy --> id is required
			secure.DELETE("/policies/:id", controllers.DeletePolicy)

			/* Credentials routes */

			/* GitHub credentials routes */

			// Find all GitHub credentials
			secure.GET("/credentials/github", controllers.FindGitHubCredentials)

			// Find a GitHub credential by id
			secure.GET("/credentials/github/:id", controllers.FindGitHubCredentialById)

			// Create a GitHub credential
			secure.POST("/credentials/github", controllers.CreateGitHubCredential)

			// Update a GitHub credential
			secure.PATCH("/credentials/github/:id", controllers.UpdateGitHubCredential)

			// Delete a GitHub credential
			secure.DELETE("/credentials/github/:id", controllers.DeleteGitHubCredential)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
