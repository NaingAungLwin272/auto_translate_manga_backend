package services

import (
	dtos "auto_translate_manga_backend/internal/dto"
	"auto_translate_manga_backend/internal/models"
	"auto_translate_manga_backend/internal/repositories"
	"context"
)

type DownloadService struct {
	repo *repositories.DownloadRepository
}

func NewDownloadService(repo *repositories.DownloadRepository) *DownloadService {
	return &DownloadService{repo: repo}
}

func (downloadService *DownloadService) CreateDownloadMangaByChapterId(ctx context.Context, download *dtos.CreateDownloadMangaDtoById) (*models.Download, error) {
	downloadResponse := &models.Download{
		User:    download.User,
		Chapter: download.Chapter,
		Manga:   download.Manga,
		Storage: download.Storage,
	}
	return downloadService.repo.CreateDownloadMangaByChapterId(ctx, downloadResponse)
}
