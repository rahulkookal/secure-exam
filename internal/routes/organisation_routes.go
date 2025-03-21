package routes

import (
	"github.com/gin-gonic/gin"
	"rahul.com/secure-exam/internal/handlers"
)

func RegisterOrganisationRoutes(server *gin.Engine) {
	orgRoutes := server.Group("/organisation")
	{
		orgRoutes.POST("/", handlers.CreateOrganisation) // Create an organisation
		orgRoutes.GET("/all", handlers.GetOrganisations) // Get all organisations
		// orgRoutes.GET("/:id", handlers.GetOrganisationByID)   // Get organisation by ID
		// orgRoutes.PUT("/:id", handlers.UpdateOrganisation)    // Update an organisation
		// orgRoutes.DELETE("/:id", handlers.DeleteOrganisation) // Delete an organisation
	}
}
