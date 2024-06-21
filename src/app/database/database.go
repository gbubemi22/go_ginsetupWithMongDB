package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() (*mongo.Client, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("failed to load environment variables: %w", err)
	}

	MongoDb := os.Getenv("MONGO_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		return nil, fmt.Errorf("failed to create MongoDB client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	fmt.Println("Connected to MongoDB")
	return client, nil
}

// OpenCollection opens a MongoDB collection
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("GO_Rental").Collection(collectionName)
	return collection
}
