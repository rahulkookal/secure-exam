package repository

import (
	"context"
	"time"

	"github.com/go-errors/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (repo *MongoRepository) GetUserByRegistrationID(registrationID string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"registration_id": registrationID, "is_deleted": false}
	var user model.User
	err := repo.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // No user found
		}
		return nil, err
	}

	return &user, nil
}

func (repo *MongoRepository) GetStudentsByOrganisation(orgID primitive.ObjectID) ([]model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"role":            "STUDENT",
		"organisation_id": orgID,
		"is_deleted":      false,
	}

	var students []model.User
	cursor, err := repo.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &students); err != nil {
		return nil, err
	}

	return students, nil
}
