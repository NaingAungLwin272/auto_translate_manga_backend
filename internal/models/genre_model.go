package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Genre struct {
	ID          bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
}
