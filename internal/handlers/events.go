package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rahul.com/secure-exam/internal/model"
	"rahul.com/secure-exam/internal/repository"
)

var (
	eventRepo = repository.NewMongoRepository("secure-exam", "event")
)

func CreateEvent(ctx *gin.Context) {
	var eventQuery model.EventQuery
	if err := ctx.ShouldBindJSON(&eventQuery); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	userID, err := getUserID(eventQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingEvent, err := eventRepo.GetEvent(eventQuery.ExamID, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking for existing event"})
		return
	}
	if existingEvent != nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Event already exists", "data": existingEvent})
		return
	}

	newEvent := model.Event{
		ID:        primitive.NewObjectID(),
		ExamID:    eventQuery.ExamID,
		UserID:    userID,
		Status:    "SCHEDULED",
		Score:     0,
		StartedAt: time.Now(),
		IsDeleted: false,
		UpdatedAt: time.Now(),
	}

	if err := eventRepo.CreateEvent(&newEvent); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": newEvent})
}

func StartEvent(ctx *gin.Context) {
	var eventQuery model.EventUpdateQuery
	if err := ctx.ShouldBindJSON(&eventQuery); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	event, err := eventRepo.GetEventById(eventQuery.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching event"})
		return
	}
	if event == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	questions, err := questionRepo.FindQuestionsByExamID(event.ExamID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching questions"})
		return
	}
	fmt.Print(questions)
	if err := eventRepo.UpdateEventQuestions(eventQuery.ID, questions); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating event questions"})
		return
	}
	event.Questions = questions

	ctx.JSON(http.StatusOK, gin.H{"data": event})
}

func UpdateEventQuestionAnswers(ctx *gin.Context) {
	var req struct {
		EventID    string   `json:"event_id" binding:"required"`
		QuestionID string   `json:"question_id" binding:"required"`
		Answers    []string `json:"answers"`
	}

	// Bind request JSON body to struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	// Convert EventID to ObjectID
	eventID, err := primitive.ObjectIDFromHex(req.EventID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event_id", "details": err.Error()})
		return
	}

	// Convert QuestionID to ObjectID
	questionID, err := primitive.ObjectIDFromHex(req.QuestionID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question_id", "details": err.Error()})
		return
	}

	// Call repository function to update answers
	fmt.Println("========")
	fmt.Println(req.Answers)
	err = eventRepo.UpdateEventQuestionAnswers(eventID, questionID, req.Answers)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update answers", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Question answers updated successfully"})
}

func getUserID(eventQuery model.EventQuery) (primitive.ObjectID, error) {
	if !eventQuery.UserID.IsZero() {
		return eventQuery.UserID, nil
	}
	if eventQuery.RegistrationID != "" {
		user, err := userRepo.GetUserByRegistrationID(eventQuery.RegistrationID)
		if err != nil {
			return primitive.NilObjectID, err
		}
		if user == nil {
			return primitive.NilObjectID, fmt.Errorf("user not found")
		}
		return user.ID, nil
	}
	return primitive.NilObjectID, fmt.Errorf("either UserID or RegistrationID must be provided")
}
