package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Genre struct {
	ID          bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description,omitempty" json:"description,omitempty"`
	CreatedAt   time.Time     `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt   time.Time     `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
