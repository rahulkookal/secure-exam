package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"rahul.com/secure-exam/config"
)

// MongoRepository implements CRUD for MongoDB
type MongoRepository struct {
	collection *mongo.Collection
}

// NewMongoRepository creates a new repository
func NewMongoRepository(dbName, collectionName string) *MongoRepository {
	client := config.ConnectMongo()
	collection := client.Database(dbName).Collection(collectionName)
	return &MongoRepository{collection}
}
