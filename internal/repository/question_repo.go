package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rahul.com/secure-exam/internal/model"
)

func (repo *MongoRepository) CreateQuestion(question *model.Question) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	question.ID = primitive.NewObjectID()
	_, err := repo.collection.InsertOne(ctx, question)
	return err
}

func (repo *MongoRepository) UpdateQuestion(question *model.Question) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": question.ID} // Match by question ID
	update := bson.M{"$set": bson.M{
		"question_text":   question.QuestionText,
		"options":         question.Options,
		"correct_answers": question.CorrectAnswers,
		"is_multi_select": question.IsMultiSelect,
		"is_deleted":      question.IsDeleted,
	}}

	_, err := repo.collection.UpdateOne(ctx, filter, update)
	return err
}
