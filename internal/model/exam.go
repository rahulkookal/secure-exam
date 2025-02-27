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
	ValidFrom   time.Time          `bson:"valid_from" json:"valid_from" binding:"required"`
	ValidTill   time.Time          `bson:"valid_till" json:"valid_till" binding:"required"`
	MaxQuestion int                `bson:"max_question" json:"max_question" binding:"required"`
}
