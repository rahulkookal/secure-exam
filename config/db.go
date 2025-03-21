package config

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var once sync.Once

func ConnectMongo() *mongo.Client {
	once.Do(func() {
		// Fetch environment variables
		mongoURL := os.Getenv("MONGO_URL")
		fmt.Println("MongoDB URL:", mongoURL)

		// Default value if MONGO_URL is not set
		if mongoURL == "" {
			mongoURL = "mongodb+srv://secure-admin:Admin123%21@secureexam.v43yx.mongodb.net/?retryWrites=true&w=majority&appName=SecureExam"
		}
		fmt.Println("MongoDB URL:", mongoURL)
		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		opts := options.Client().ApplyURI(mongoURL).SetServerAPIOptions(serverAPI)

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
