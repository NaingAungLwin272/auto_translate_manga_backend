package dtos

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type CreateFavoritedDTO struct {
	User  bson.ObjectID `bson:"user" json:"user"`
	Manga bson.ObjectID `bson:"manga" json:"manga"`
}

type FavoritedResponseDTO struct {
	ID        bson.ObjectID `json:"id"`
	User      bson.ObjectID `json:"user"`
	Manga     bson.ObjectID `json:"manga"`
	CreatedAt time.Time     `json:"created_at"`
	DeletedAt time.Time     `json:"deleted_at"`
}

