package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Exam struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title             string             `bson:"title" json:"title" binding:"required"`
	Description       string             `bson:"description" json:"description" binding:"required"`
	Duration          int                `bson:"duration" json:"duration" binding:"required"` // In minutes
	ValidFrom         time.Time          `bson:"valid_from" json:"valid_from" binding:"required"`
	ValidTill         time.Time          `bson:"valid_till" json:"valid_till" binding:"required"`
	NumberOfQuestions int                `bson:"number_of_questions" json:"number_of_questions" binding:"required"`
	OrganisationID    primitive.ObjectID `bson:"organisation_id" json:"organisation_id" binding:"required"`
	IsDeleted         bool               `bson:"is_deleted" json:"is_deleted"`
}
