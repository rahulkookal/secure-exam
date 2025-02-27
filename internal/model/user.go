package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName    string             `bson:"first_name" json:"first_name" binding:"required"`
	LastName     string             `bson:"last_name" json:"last_name" binding:"required"`
	Email        string             `bson:"email" json:"email" binding:"required,email"`
	PasswordHash string             `bson:"password_hash" json:"password_hash"`
	Password     string             `json:"password,omitempty" binding:"required"`
	Verified     bool               `bson:"verified" json:"verified"`
	Mobile       string             `bson:"mobile" json:"mobile"`
	Token        string             `bson:"token" json:"token,omitempty"`
	IsAdmin      bool               `bson:"is_admin" json:"is_admin"`
}
