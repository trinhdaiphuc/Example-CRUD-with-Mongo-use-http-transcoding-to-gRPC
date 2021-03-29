package services

import "go.mongodb.org/mongo-driver/mongo"

type Entities struct {
	EntityCollection *mongo.Collection
}
