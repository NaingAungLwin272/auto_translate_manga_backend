package dtos

import (
	"auto_translate_manga_backend/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type CreateMangaDTO struct {
	Title       string          `json:"title" binding:"required"`
	Description string          `json:"description,omitempty"`
	Author      string          `json:"author,omitempty"`
	CoverImage  string          `json:"cover_image" binding:"required"`
	Genres      []bson.ObjectID `json:"genres" binding:"required"` // references genre ids
	Status      string          `json:"status" binding:"required"` // "ongoing" or "completed"
}

type MangaResponseDTO struct {
	ID          string          `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description,omitempty"`
	Author      string          `json:"author,omitempty"`
	CoverImage  string          `json:"cover_image"`
	Genres      []bson.ObjectID `json:"genres"` // references genre ids
	Status      string          `json:"status"` // "ongoing" or "completed"
	CreatedAt   time.Time       `json:"created_at,omitempty"`
	UpdatedAt   time.Time       `json:"updated_at,omitempty"`
}

type MangaResponseDTOWithGenre struct {
	ID          string         `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description,omitempty"`
	Author      string         `json:"author,omitempty"`
	CoverImage  string         `json:"cover_image"`
	Genres      []models.Genre `json:"genres"` // references genre ids
	Status      string         `json:"status"` // "ongoing" or "completed"
	CreatedAt   time.Time      `json:"created_at,omitempty"`
	UpdatedAt   time.Time      `json:"updated_at,omitempty"`
}
