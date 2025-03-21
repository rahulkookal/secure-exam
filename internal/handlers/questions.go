package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rahul.com/secure-exam/internal/model"
	"rahul.com/secure-exam/internal/repository"
)

var questionRepo = repository.NewMongoRepository("secure-exam", "questions")

func CreateQuestion(ctx *gin.Context) {
	var question model.Question
	err := ctx.ShouldBindJSON(&question)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request", "error": err})
		return
	}

	if err := questionRepo.CreateQuestion(&question); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": question})
}

func UpdateQuestion(ctx *gin.Context) {
	// Get question ID from URL params
	questionID := ctx.Param("id")
	if questionID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Question ID is required"})
		return
	}

	// Convert questionID string to primitive.ObjectID
	objID, err := primitive.ObjectIDFromHex(questionID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Question ID format"})
		return
	}

	// Bind request body to question model
	var question model.Question
	if err := ctx.ShouldBindJSON(&question); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request", "error": err.Error()})
		return
	}

	// Ensure the ID in request matches the one in the path
	question.ID = objID

	// Call repository function to update the question
	if err := questionRepo.UpdateQuestion(&question); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Question updated successfully", "data": question})
}
