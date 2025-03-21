package routes

import (
	"github.com/gin-gonic/gin"
	"rahul.com/secure-exam/internal/auth"
	"rahul.com/secure-exam/internal/handlers"
)

func RegisterUserRoutes(server *gin.Engine) {
	userRoutes := server.Group("/user")
	{
		userRoutes.POST("/register", handlers.RegisterUser)
		userRoutes.POST("/login", handlers.LoginUser)
		userRoutes.GET("/students", auth.AuthMiddleware(), auth.AdminPrivilage(), handlers.GetStudents)
	}
}
