package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Favorite struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	User      bson.ObjectID `bson:"user" json:"user"`
	Manga     bson.ObjectID `bson:"manga" json:"manga"`
	CreatedAt time.Time     `bson:"created_at,omitempty" json:"created_at,omitempty"`
	DeletedAt time.Time     `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

type FavoriteWithUsersMangaDetail struct {
	User      User           `bson:"user" json:"user"`
	Favorites []FavoriteItem `bson:"favorites" json:"favorites"`
}

type FavoriteItem struct {
	Manga     MangaWithGenres `bson:"manga" json:"manga"`
	CreatedAt time.Time       `bson:"created_at" json:"created_at"`
}
