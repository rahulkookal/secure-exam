package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ExamID    primitive.ObjectID `bson:"exam_id" json:"exam_id" binding:"required"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id" binding:"required"`
	Status    string             `bson:"status" json:"status" binding:"required,oneof=SCHEDULED STARTED COMPLETED"`
	Score     int                `bson:"score" json:"score"`
	StartedAt time.Time          `bson:"started_at" json:"started_at"`
	Questions []Question         `bson:"questions,omitempty" json:"questions,omitempty"`
	IsDeleted bool               `bson:"is_deleted" json:"is_deleted"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type EventQuery struct {
	ExamID         primitive.ObjectID `bson:"exam_id" json:"exam_id" binding:"required"`
	UserID         primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	RegistrationID string             `bson:"registration_id,omitempty" json:"registration_id,omitempty"`
}
type EventUpdateQuery struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
}

type EventAnswers struct {
	EventID    primitive.ObjectID `bson:"event_id" json:"event_id" binding:"required"`
	QuestionID primitive.ObjectID `bson:"question_id" json:"question_id" binding:"required"`
	Answers    []string           `bson:"answers" json:"answers"`
}
