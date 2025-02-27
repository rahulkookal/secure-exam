package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"rahul.com/secure-exam/internal/auth"
	"rahul.com/secure-exam/internal/model"
	"rahul.com/secure-exam/internal/repository"
)

var userRepo = repository.NewMongoRepository("secure-exam", "users")

func RegisterUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if error := userRepo.FoundExcistingUser(user.Email); error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.PasswordHash = string(hashedPassword)
	user.Verified = false
	user.Password = ""

	if err := userRepo.CreateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": user})
}

func LoginUser(ctx *gin.Context) {
	var cred struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&cred); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}
	user, err := userRepo.FindUser(cred.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(cred.Password))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	token, _ := auth.GenerateJWT(user.ID.Hex(), user.Email, user.IsAdmin)
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
