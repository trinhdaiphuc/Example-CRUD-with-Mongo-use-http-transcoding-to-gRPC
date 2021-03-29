package models

import (
	"github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"
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
