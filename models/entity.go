package models

import (
	"context"
	"fmt"
	"os"

	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EntityItem struct {
	ID          uuid.UUID `bson:"_id,omitempty"`
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
	URL         string    `bson:"url"`
}

func NewEntityCollection(db *mongo.Client) (entityCollection *mongo.Collection) {
	entityCollection = db.Database("mydb").Collection("entity")
	return
}

func InitDB() (*mongo.Client, context.Context, error) {
	// Initialize MongoDb client
	fmt.Println("Connecting to MongoDB...")

	mongoURI := "mongodb://localhost:27017"
	if len(os.Getenv("DB_HOST")) > 0 {
		mongoURI = os.Getenv("DB_HOST")
	}

	// non-nil empty context
	mongoCtx := context.Background()

	// Connect takes in a context and options, the connection URI is the only option we pass for now
	db, err := mongo.Connect(mongoCtx, options.Client().ApplyURI(mongoURI))

	// Handle potential errors
	if err != nil {
		return nil, mongoCtx, err
	}

	// Check whether the connection was successful by pinging the MongoDB server
	err = db.Ping(mongoCtx, nil)
	if err != nil {
		return nil, mongoCtx, err
	}
	fmt.Println("Connected to Mongodb")
	return db, mongoCtx, nil
}
