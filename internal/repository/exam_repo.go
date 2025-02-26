package repository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rahul.com/secure-exam/internal/model"
)

// Create a new Exam
func (repo *MongoRepository) Create(exam *model.Exam) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	exam.ID = primitive.NewObjectID()
	_, err := repo.collection.InsertOne(ctx, exam)
	return err
}

func (repo *MongoRepository) GetExams() ([]model.Exam, error) {
	var exams []model.Exam
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := repo.collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("MongoDB Find() error:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	// Debugging: Check if cursor has results
	if !cursor.Next(ctx) {
		log.Println("No exams found in MongoDB")
		return []model.Exam{}, nil
	}

	if err := cursor.All(ctx, &exams); err != nil {
		log.Println("Cursor decoding error:", err)
		return nil, err
	}

	log.Printf("Fetched %d exams", len(exams))
	return exams, nil
}

func (repo *MongoRepository) GetExamByID(id primitive.ObjectID) (*model.Exam, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var exam model.Exam
	err := repo.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&exam)
	if err != nil {
		return nil, err
	}
	return &exam, nil
}
