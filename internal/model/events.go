package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID        primitive.ObjectID            `bson:"_id,omitempty" json:"id"`
	ExamID    primitive.ObjectID            `bson:"exam_id" json:"exam_id" binding:"required"`
	UserID    primitive.ObjectID            `bson:"user_id" json:"user_id" binding:"required"`
	Status    string                        `bson:"status" json:"status" binding:"required,oneof=SCHEDULED STARTED COMPLETED"`
	Score     int                           `bson:"score" json:"score"`
	StartedAt time.Time                     `bson:"started_at" json:"started_at"`
	Answers   map[primitive.ObjectID]int    `bson:"answers,omitempty" json:"answers,omitempty"`
	Attempts  map[primitive.ObjectID]string `bson:"attempts,omitempty" json:"attempts,omitempty"`
}
