package controllers

import (
	dtos "auto_translate_manga_backend/internal/dto"
	"auto_translate_manga_backend/internal/services"
	"auto_translate_manga_backend/internal/utils"
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ChapterController struct {
	service *services.ChapterService
}

func NewChapterController(service *services.ChapterService) *ChapterController {
	return &ChapterController{service: service}
}

func (chapterController *ChapterController) CreateChapter(ctx *gin.Context) {
	var dto dtos.CreateChapterDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   http.StatusText(http.StatusBadRequest),
			Message: "fields are require",
		})
		return
	}

	chapter, err := chapterController.service.CreateChapter(context.Background(), &dto)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			ctx.JSON(http.StatusConflict, utils.ErrorResponse{
				Status:  http.StatusConflict,
				Error:   http.StatusText(http.StatusConflict),
				Message: "chapter already exists",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   http.StatusText(http.StatusInternalServerError),
			Message: "internal server error",
		})
		return
	}

	chapterResponse := dtos.ChapterResponseDTO{
		ID:            chapter.ID,
		Manga:         chapter.Manga,
		Title:         chapter.Title,
		ChapterNumber: chapter.ChapterNumber,
		Pages:         chapter.Pages,
		CreatedAt:     chapter.CreatedAt,
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse[dtos.ChapterResponseDTO]{
		Status:  http.StatusOK,
		Message: "chapter created successfully",
		Data:    chapterResponse,
	})
}

func (chapterController *ChapterController) GetAllChaptersByMangaID(ctx *gin.Context) {
	id := ctx.Params.ByName("manga_id")
	manga, err := chapterController.service.GetAllChaptersByMangaID(ctx, id)
	if err != nil {
		if err.Error() == "manga not found" {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
				Status:  http.StatusNotFound,
				Error:   http.StatusText(http.StatusNotFound),
				Message: "manga not found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
		return
	}

	var chapterResponse []dtos.ChapterResponseDTOWithMangaGenres
	for _, result := range *manga {
		chapterResponse = append(chapterResponse, dtos.ChapterResponseDTOWithMangaGenres{
			ID:            result.ID.Hex(),
			Manga:         result.Manga,
			ChapterNumber: result.ChapterNumber,
			Title:         result.Title,
			Pages:         result.Pages,
			CreatedAt:     result.CreatedAt,
			UpdatedAt:     result.UpdatedAt,
		})
	}

	if len(chapterResponse) == 0 {
		ctx.JSON(http.StatusOK, utils.SuccessResponse[[]dtos.ChapterResponseDTOWithMangaGenres]{
			Status:  http.StatusOK,
			Message: "no chapter found",
			Data:    []dtos.ChapterResponseDTOWithMangaGenres{},
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse[[]dtos.ChapterResponseDTOWithMangaGenres]{
		Status:  http.StatusOK,
		Message: "chapter found successfully",
		Data:    chapterResponse,
	})
}
