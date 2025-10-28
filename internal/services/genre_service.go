package services

import (
	dtos "auto_translate_manga_backend/internal/dto"
	"auto_translate_manga_backend/internal/models"
	"auto_translate_manga_backend/internal/repositories"
	"context"
)

type GenreService struct {
	repo *repositories.GenreRepo
}

func NewGenreService(repo *repositories.GenreRepo) *GenreService {
	return &GenreService{repo: repo}
}

func (genreService *GenreService) CreateGenre(ctx context.Context, dto *dtos.CreateGenreDTO) (*models.Genre, error) {
	genre := &models.Genre{
		Name:        dto.Name,
		Description: dto.Description,
	}
	return genreService.repo.CreateGenre(ctx, genre)
}

func (genreService *GenreService) GetAllGenres(ctx context.Context) ([]models.Genre, error) {
	return genreService.repo.GetAllGenres(ctx)
}

func (genreService *GenreService) GetGenreById(ctx context.Context, id string) (*models.Genre, error) {
	return genreService.repo.GetGenreById(ctx, id)
}

func (genreService *GenreService) UpdateGenreById(ctx context.Context, id string, genre *dtos.UpdateGenreDTO) (*models.Genre, error) {
	updatedGenre, err := genreService.repo.UpdateGenreById(ctx, id, genre)
	if err != nil {
		return nil, err
	}
	return updatedGenre, nil
}

func (genreService *GenreService) DeleteGenreById(ctx context.Context, id string) (string, error) {
	result, err := genreService.repo.DeleteGenreById(ctx, id)
	if err != nil {
		return "", err
	}
	return result, nil
}
