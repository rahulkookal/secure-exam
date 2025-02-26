package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Exam struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title" binding:"required"`
	Description string             `bson:"description" json:"description" binding:"required"`
	Duration    int                `bson:"duration" json:"duration" binding:"required"` // In minutes
	StartTime   time.Time          `bson:"start_time" json:"start_time" binding:"required"`
	EndTime     time.Time          `bson:"end_time" json:"end_time" binding:"required"`
	MaxQuestion int                `bson:"max_question" json:"max_question" binding:"required"`
}
