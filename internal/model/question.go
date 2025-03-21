package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Question struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ExamID         primitive.ObjectID `bson:"exam_id" json:"exam_id" binding:"required"` // Reference to Exam
	QuestionText   string             `bson:"question_text" json:"question_text" binding:"required"`
	Options        []string           `bson:"options" json:"options" binding:"required"`
	CorrectAnswers []string           `bson:"correct_answers" json:"correct_answers" binding:"required"` // Changed to string array
	IsMultiSelect  bool               `bson:"is_multi_select" json:"is_multi_select"`
	IsDeleted      bool               `bson:"is_deleted" json:"is_deleted"`
}
