package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EntityItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	URL         string             `bson:"url"`
}

func NewEntityCollection(db *mongo.Client) (entityCollection *mongo.Collection) {
	entityCollection = db.Database("mydb").Collection("entity")
	return
}
