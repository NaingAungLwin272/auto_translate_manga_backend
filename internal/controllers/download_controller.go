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

type DownloadController struct {
	service        *services.DownloadService
	chapterService *services.ChapterService
}

func NewDownloadController(service *services.DownloadService, chapterService *services.ChapterService) *DownloadController {
	return &DownloadController{service: service, chapterService: chapterService}
}

func (downloadController *DownloadController) CreateDownloadManga(ctx *gin.Context) {
	var dto dtos.CreateDownloadMangaDtoById
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   http.StatusText(http.StatusBadRequest),
			Message: "fields are require",
		})
		return
	}
	id := dto.Chapter.Hex()

	chapterResult, err := downloadController.chapterService.GetChapterByID(context.Background(), id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
				Status:  http.StatusNotFound,
				Error:   http.StatusText(http.StatusNotFound),
				Message: "chapter not found",
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

	dto.Chapter = chapterResult.ID
	dto.Manga = chapterResult.Manga.ID

	download, err := downloadController.service.CreateDownloadMangaByChapterId(context.Background(), &dto)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			ctx.JSON(http.StatusConflict, utils.ErrorResponse{
				Status:  http.StatusConflict,
				Error:   http.StatusText(http.StatusConflict),
				Message: "download already exists",
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

	downloadResponse := dtos.DownloadResponseDTO{
		ID:        download.ID,
		User:      download.User,
		Chapter:   download.Chapter,
		Manga:     download.Manga,
		Storage:   download.Storage,
		CreatedAt: download.CreatedAt,
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse[dtos.DownloadResponseDTO]{
		Status:  http.StatusOK,
		Message: "download successfully",
		Data:    downloadResponse,
	})
}
