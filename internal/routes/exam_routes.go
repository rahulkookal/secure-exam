package routes

import (
	"github.com/gin-gonic/gin"
	"rahul.com/secure-exam/internal/auth"
	"rahul.com/secure-exam/internal/handlers"
)

func RegisterExamRoutes(server *gin.Engine) {
	examRoutes := server.Group("/exam")
	{
		examRoutes.GET("/all", auth.AuthMiddleware(), handlers.GetExams)
		examRoutes.GET("/", auth.AuthMiddleware(), handlers.GetExamByID)
		examRoutes.POST("/", auth.AuthMiddleware(), handlers.CreateExams)
		examRoutes.POST("/question", auth.AuthMiddleware(), auth.AdminPrivilage(), handlers.CreateQuestion)
		examRoutes.PUT("/question/:id", auth.AuthMiddleware(), auth.AdminPrivilage(), handlers.UpdateQuestion)
	}
}
