package utils

import (
	"os"
	"fmt"
	"log"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoClient(ctx context.Context) (*mongo.Client, error) {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://root:admin@127.0.0.1:27017"
	}
	fmt.Printf("mongoURI: %v\n", mongoURI)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return client, nil
}

func GetMongoCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("gogrpc").Collection(collectionName)
}
