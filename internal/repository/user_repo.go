package repository

import (
	"context"
	"time"

	"github.com/go-errors/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rahul.com/secure-exam/internal/model"
)

// Create a new Exam
func (repo *MongoRepository) CreateUser(user model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user.ID = primitive.NewObjectID()
	_, err := repo.collection.InsertOne(ctx, user)
	return err
}

func (repo *MongoRepository) FoundExcistingUser(email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var existingUser model.User
	err := repo.collection.FindOne(ctx, bson.M{"email": email}).Decode(&existingUser)

	if err == nil {
		return errors.New("email already registered")
	}

	return nil // Other database error
}

func (repo *MongoRepository) FindUser(email string) (model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user model.User
	err := repo.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	return user, err
}
