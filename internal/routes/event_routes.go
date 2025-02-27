package routes

import (
	"github.com/gin-gonic/gin"
	"rahul.com/secure-exam/internal/auth"
	"rahul.com/secure-exam/internal/handlers"
)

func RegisterEventRoutes(server *gin.Engine) {
	examRoutes := server.Group("/event")
	{
		examRoutes.POST("/", auth.AuthMiddleware(), handlers.CreateExams)
	}
}
