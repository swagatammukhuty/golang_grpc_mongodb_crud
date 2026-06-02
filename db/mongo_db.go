package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserCollection *mongo.Collection

func InitMongo() error {
	//Provide cancel context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Create the URL
	clientOptions := options.Client().ApplyURI("mongodb://localhost://27017")
	// Connect the mongodb with URL and proper cancellation
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	UserCollection = client.Database("grpc").Collection("users")
	return nil
}
