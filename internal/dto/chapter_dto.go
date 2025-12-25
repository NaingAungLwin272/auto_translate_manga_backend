package dtos

import (
	"auto_translate_manga_backend/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type CreateChapterDTO struct {
	Manga         bson.ObjectID          `json:"manga"`
	ChapterNumber int64                  `json:"chapter_number"`
	Title         string                 `json:"title"`
	Pages         []models.ChapterDetail `json:"pages"`
}

type ChapterResponseDTO struct {
	ID            bson.ObjectID          `json:"id"`
	Manga         bson.ObjectID          `json:"manga"`
	ChapterNumber int64                  `json:"chapter_number"`
	Title         string                 `json:"title"`
	Pages         []models.ChapterDetail `json:"pages"`
	CreatedAt     time.Time              `json:"created_at"`
}

type ChapterResponseDTOWithMangaGenres struct {
	ID            string                 `json:"id"`
	Manga         models.MangaWithGenres `json:"manga"`
	ChapterNumber int64                  `json:"chapter_number"`
	Title         string                 `json:"title"`
	Pages         []models.ChapterDetail `json:"pages"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
}
