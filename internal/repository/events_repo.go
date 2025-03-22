package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
func (repo *MongoRepository) GetEventById(eventID primitive.ObjectID) (*model.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": eventID, "is_deleted": false}
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

func (repo *MongoRepository) UpdateEventQuestions(eventID primitive.ObjectID, questions []model.Question) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": eventID}
	update := bson.M{"$set": bson.M{"questions": questions, "updated_at": time.Now()}}

	_, err := repo.collection.UpdateOne(ctx, filter, update)
	return err
}

func (repo *MongoRepository) UpdateEventQuestionAnswers(eventID, questionID primitive.ObjectID, answers []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Ensure answers is at least an empty array (avoid setting null)
	if answers == nil {
		answers = []string{}
	}

	// Filter to find the event document
	filter := bson.M{
		"_id":        eventID,
		"is_deleted": false,
	}

	// Update query to modify only the `answers` field inside the matching question
	update := bson.M{
		"$set": bson.M{
			"questions.$[elem].answers": answers,
		},
	}

	// Array filter to target the specific question inside `questions`
	updateOptions := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{
			bson.M{"elem._id": questionID}, // ✅ Ensure `_id` matches `questionID`
		},
	})
	log.Printf("Updating answers for eventID: %s, questionID: %s", eventID.Hex(), questionID.Hex())
	log.Printf("Updating with answers: %+v", answers)
	// Execute the update
	result, err := repo.collection.UpdateOne(ctx, filter, update, updateOptions)
	if err != nil {
		return fmt.Errorf("MongoDB update error: %v", err)
	}

	// Debugging output
	if result.MatchedCount == 0 {
		return fmt.Errorf("no matching event found for eventID: %s", eventID.Hex())
	}
	if result.ModifiedCount == 0 {
		return fmt.Errorf("event found, but question with _id: %s was not updated", questionID.Hex())
	}

	log.Printf("Successfully updated answers for question %s in event %s", questionID.Hex(), eventID.Hex())
	return nil
}
