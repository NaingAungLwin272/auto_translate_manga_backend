package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Manga struct {
	ID          bson.ObjectID   `bson:"_id,omitempty" json:"_id"`
	Title       string          `bson:"title" json:"title"`
	Description string          `bson:"description,omitempty" json:"description,omitempty"`
	Author      string          `bson:"author,omitempty" json:"author,omitempty"`
	CoverImage  string          `bson:"cover_image" json:"cover_image"`
	Genres      []bson.ObjectID `bson:"genres,omitempty" json:"genres,omitempty"` // references genre ids
	Status      string          `bson:"status,omitempty" json:"status,omitempty"` // "ongoing" or "completed"
	CreatedAt   time.Time       `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt   time.Time       `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type MangaWithGenres struct {
	ID          bson.ObjectID `bson:"_id" json:"_id"`
	Title       string        `bson:"title" json:"title"`
	Description string        `bson:"description" json:"description"`
	Author      string        `bson:"author" json:"author"`
	CoverImage  string        `bson:"cover_image" json:"cover_image"`
	Genres      []Genre       `bson:"genres" json:"genres"`
	Status      string        `bson:"status" json:"status"`
	CreatedAt   time.Time     `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt   time.Time     `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
