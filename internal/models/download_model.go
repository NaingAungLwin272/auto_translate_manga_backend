package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Storage struct {
	Path string `bson:"path" json:"path"`
	Size int64  `bson:"size" json:"size"`
}

type Download struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	User      bson.ObjectID `bson:"user" json:"user"`
	Chapter   bson.ObjectID `bson:"chapter" json:"chapter"`
	Manga     bson.ObjectID `bson:"manga" json:"manga"`
	Storage   Storage       `bson:"storage" json:"storage"`
	CreatedAt time.Time     `bson:"created_at,omitempty" json:"created_at,omitempty"`
	DeletedAt time.Time     `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}
