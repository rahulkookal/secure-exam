package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
