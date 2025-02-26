package repository

import (
	"context"
	"time"

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
