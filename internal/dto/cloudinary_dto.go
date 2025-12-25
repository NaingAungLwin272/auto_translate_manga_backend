package dtos

import "auto_translate_manga_backend/internal/models"

type UploadImageResponse struct {
	File_Url string `json:"file_url"`
}

type UploadPagesResponse struct {
	Pages []models.ChapterDetail `json:"pages"`
}

