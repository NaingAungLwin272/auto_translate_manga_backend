package services

import (
	dtos "auto_translate_manga_backend/internal/dto"
	"auto_translate_manga_backend/internal/enum"
	"auto_translate_manga_backend/internal/models"
	"auto_translate_manga_backend/internal/repositories"
	"context"
	"fmt"
)

type MangaService struct {
	repo *repositories.MangaRepo
}

func NewMangaService(repo *repositories.MangaRepo) *MangaService {
	return &MangaService{repo: repo}
}

func (mangaService *MangaService) CreateManga(ctx context.Context, mangaDTO *dtos.CreateMangaDTO) (*models.Manga, error) {
	if !enum.MangaStatus(mangaDTO.Status).IsValid() {
		return nil, fmt.Errorf("invalid status: %s", mangaDTO.Status)
	}
	manga := &models.Manga{
		Title:       mangaDTO.Title,
		Description: mangaDTO.Description,
		Author:      mangaDTO.Author,
		CoverImage:  mangaDTO.CoverImage,
		Genres:      mangaDTO.Genres,
		Status:      mangaDTO.Status,
	}
	return mangaService.repo.CreateManga(ctx, manga)
}

func (mangaService *MangaService) GetAllManga(ctx context.Context) ([]models.MangaWithGenres, error) {
	return mangaService.repo.GetAllManga(ctx)
}

func (mangaService *MangaService) GetMangaById(ctx context.Context, id string) (*models.MangaWithGenres, error) {
	return mangaService.repo.GetMangaById(ctx, id)
}

func (mangaService *MangaService) FilterManga(ctx context.Context, title string, genres []string, status string) (*[]models.MangaWithGenres, error) {
	return mangaService.repo.FilterManga(ctx, title, genres, status)
}

func (mangaService *MangaService) DeleteManga(ctx context.Context, id string) (string, error) {
	return mangaService.repo.DeleteManga(ctx, id)
}
