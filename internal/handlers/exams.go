package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rahul.com/secure-exam/internal/model"
	"rahul.com/secure-exam/internal/repository"
)

var examRepo = repository.NewMongoRepository("secure-exam", "exams")

func GetExams(ctx *gin.Context) {
	exams, err := examRepo.GetExams()
	if err != nil {
		log.Println("Error fetching exams:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch exams"})
		return
	}

	if exams == nil {
		exams = []model.Exam{} // Ensuring empty array instead of null
	}

	ctx.JSON(http.StatusOK, gin.H{"data": exams})
}

func CreateExams(ctx *gin.Context) {
	var exam model.Exam
	err := ctx.ShouldBindJSON(&exam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request", "error": err})
		return
	}
	if err := examRepo.Create(&exam); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": exam})
}

func GetExamByID(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing exam ID"})
		return
	}

	// Convert string ID to MongoDB ObjectID
	examID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid exam ID format"})
		return
	}

	exam, err := examRepo.GetExamByID(examID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Exam not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"exam": exam})
}
