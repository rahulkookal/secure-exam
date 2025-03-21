package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Organisation struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name" binding:"required"`
	Email     string             `bson:"email" json:"email" binding:"required,email"`
	Phone     string             `bson:"phone" json:"phone" binding:"required"`
	Address   string             `bson:"address" json:"address" binding:"required"`
	IsDeleted bool               `bson:"is_deleted" json:"is_deleted"` // Soft delete flag
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
