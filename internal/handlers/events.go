package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rahul.com/secure-exam/internal/model"
	"rahul.com/secure-exam/internal/repository"
)

var eventRepo = repository.NewMongoRepository("secure-exam", "event")

func CreateEvent(ctx *gin.Context) {
	var eventQuery model.EventQuery
	err := ctx.ShouldBindJSON(&eventQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad request", "error": err.Error()})
		return
	}

	var userID primitive.ObjectID

	// If UserID is provided, use it directly
	if !eventQuery.UserID.IsZero() {
		userID = eventQuery.UserID
	} else if eventQuery.RegistrationID != "" {
		// Otherwise, fetch UserID using RegistrationID
		user, err := userRepo.GetUserByRegistrationID(eventQuery.RegistrationID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding user"})
			return
		}
		if user == nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		userID = user.ID
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Either UserID or RegistrationID must be provided"})
		return
	}

	// Check if an event already exists for this ExamID and UserID
	existingEvent, err := eventRepo.GetEvent(eventQuery.ExamID, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking for existing event"})
		return
	}

	if existingEvent != nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Event already exists", "data": existingEvent})
		return
	}

	// Create a new event
	newEvent := model.Event{
		ID:        primitive.NewObjectID(),
		ExamID:    eventQuery.ExamID,
		UserID:    userID,
		Status:    "SCHEDULED",
		Score:     0,
		StartedAt: time.Now(),
		Answers:   make(map[primitive.ObjectID][]string),
		Attempts:  make(map[primitive.ObjectID]string),
		IsDeleted: false,
		UpdatedAt: time.Now(),
	}

	// Save the event
	if err := eventRepo.CreateEvent(&newEvent); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": newEvent})
}
