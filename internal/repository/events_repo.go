package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"rahul.com/secure-exam/internal/model"
)

func (repo *MongoRepository) CreateEvent(event *model.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	event.ID = primitive.NewObjectID()
	_, err := repo.collection.InsertOne(ctx, event)
	return err
}

func (repo *MongoRepository) GetEvent(examID, userID primitive.ObjectID) (*model.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"exam_id": examID, "user_id": userID, "is_deleted": false}
	var event model.Event
	err := repo.collection.FindOne(ctx, filter).Decode(&event)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // No event found
		}
		return nil, err
	}

	return &event, nil
}
