package config

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var once sync.Once

func ConnectMongo() *mongo.Client {
	once.Do(func() {
		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		opts := options.Client().ApplyURI("mongodb://root:example@localhost:27017/?authSource=admin").SetServerAPIOptions(serverAPI)
		// Create a new client and connect to the server
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(ctx, opts)
		if err != nil {
			panic(err)
		}
		mongoClient = client
	})
	return mongoClient
}
