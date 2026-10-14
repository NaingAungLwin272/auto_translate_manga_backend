package dtos

import (
	"auto_translate_manga_backend/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type CreateDownloadMangaDtoById struct {
	User    bson.ObjectID  ` json:"user"`
	Chapter bson.ObjectID  ` json:"chapter"`
	Manga   bson.ObjectID  ` json:"manga"`
	Storage models.Storage ` json:"storage"`
}

type DownloadResponseDTO struct {
	ID        bson.ObjectID  `json:"id"`
	User      bson.ObjectID  ` json:"user"`
	Chapter   bson.ObjectID  ` json:"chapter"`
	Manga     bson.ObjectID  ` json:"manga"`
	Storage   models.Storage ` json:"storage"`
	CreatedAt time.Time      `json:"created_at"`
}

type DownloadResponseDTOWithMangaGenres struct {
	ID        string                 `json:"id"`
	User      bson.ObjectID          ` json:"user"`
	Chapter   bson.ObjectID          ` json:"chapter"`
	Manga     models.MangaWithGenres ` json:"manga"`
	Storage   models.Storage         ` json:"storage"`
	CreatedAt time.Time              `json:"created_at"`
	DeletedAt time.Time              `json:"deleted_at"`
}
