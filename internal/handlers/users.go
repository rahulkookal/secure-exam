package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	token, _ := auth.GenerateJWT(user.ID.Hex(), user.Email, user.OrganisationID.Hex(), user.Role)
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func GetStudents(ctx *gin.Context) {
	/// Retrieve organisation_id from context
	orgIDValue, exists := ctx.Get("organisation_id")
	if !exists {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Organisation ID not found"})
		return
	}

	// Ensure orgIDValue is a string before converting
	orgIDStr, ok := orgIDValue.(string)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Organisation ID format"})
		return
	}

	// Convert string to primitive.ObjectID
	orgID, _ := primitive.ObjectIDFromHex(orgIDStr)

	// Fetch students from the repository
	students, err := userRepo.GetStudentsByOrganisation(orgID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching students"})
		return
	}

	if students == nil {
		students = []model.User{} // Ensure empty array instead of null
	}

	ctx.JSON(http.StatusOK, gin.H{"data": students})
}
