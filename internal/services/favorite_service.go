package services

import (
	dtos "auto_translate_manga_backend/internal/dto"
	"auto_translate_manga_backend/internal/models"
	"auto_translate_manga_backend/internal/repositories"
	"context"
)

type FavoriteService struct {
	repo *repositories.FavoriteRepo
}

func NewFavoriteService(repo *repositories.FavoriteRepo) *FavoriteService {
	return &FavoriteService{repo: repo}
}

func (favoriteService *FavoriteService) CreateFavoriteManga(ctx context.Context, favorite *dtos.CreateFavoritedDTO) (*models.Favorite, error) {
	favoriteResponse := &models.Favorite{
		User:  favorite.User,
		Manga: favorite.Manga,
	}
	return favoriteService.repo.CreateFavoriteManga(ctx, favoriteResponse)
}

func (favoriteService *FavoriteService) GetFavoriteMangaByUserId(ctx context.Context, id string) (*models.FavoriteWithUsersMangaDetail, error) {
	return favoriteService.repo.GetAllFavoriteMangaByUserId(ctx, id)
}

func (favoriteService *FavoriteService) RemoveFavoriteMangaByUserId(ctx context.Context, userId string, mangaId []string) (string, error) {
	return favoriteService.repo.RemoveFavoriteMangaByUserId(ctx, userId, mangaId)
}
