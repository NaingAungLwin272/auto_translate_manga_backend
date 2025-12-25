package services

import (
	dtos "auto_translate_manga_backend/internal/dto"
	"auto_translate_manga_backend/internal/models"
	"auto_translate_manga_backend/internal/repositories"
	"context"
)

type ChapterService struct {
	repo *repositories.ChapterRepository
}

func NewChapterService(repo *repositories.ChapterRepository) *ChapterService {
	return &ChapterService{repo: repo}
}

func (chapterService *ChapterService) CreateChapter(ctx context.Context, chapter *dtos.CreateChapterDTO) (*models.Chapter, error) {
	chapterResponse := &models.Chapter{
		Manga:         chapter.Manga,
		Title:         chapter.Title,
		ChapterNumber: chapter.ChapterNumber,
		Pages:         chapter.Pages,
	}
	return chapterService.repo.CreateChapter(ctx, chapterResponse)
}

func (chapterService *ChapterService) GetAllChaptersByMangaID(ctx context.Context, mangaId string) (*[]models.ChapterWithMangaAndGenres, error) {
	return chapterService.repo.GetAllChaptersByMangaID(ctx, mangaId)
}
