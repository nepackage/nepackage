package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/otisnado/fofo-server/controllers"
	"github.com/otisnado/fofo-server/docs"
	"github.com/otisnado/fofo-server/middlewares"
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

			/* Languages routes */

			// Find all languages supported
			secure.GET("/languages", controllers.FindLanguages)

			// Find a specific language by its id
			secure.GET("languages/:id", controllers.FindLanguage)

			// Create a language
			secure.POST("/languages", controllers.CreateLanguage)

			// Update data for a language --> id is required
			secure.PATCH("/languages/:id", controllers.UpdateLanguage)

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
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
