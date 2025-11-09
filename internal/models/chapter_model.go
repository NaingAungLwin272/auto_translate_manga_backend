package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Chapter struct {
	ID            bson.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	Manga         bson.ObjectID   `bson:"manga" json:"manga"`
	ChapterNumber int64           `bson:"chapter_number" json:"chapter_number"`
	Title         string          `bson:"title" json:"title"`
	Pages         []ChapterDetail `bson:"pages" json:"pages"`
	CreatedAt     time.Time       `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt     time.Time       `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeletedAt     time.Time       `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

type ChapterDetail struct {
	PageNumber int64  `bson:"page_number" json:"page_number"`
	ImageUrl   string `bson:"image_url" json:"image_url"`
}
