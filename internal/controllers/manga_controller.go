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

type MangaController struct {
	service           *services.MangaService
	cloudinaryService *services.CloudinaryServcie
}

func NewMangaController(service *services.MangaService, cloudinaryService *services.CloudinaryServcie) *MangaController {
	return &MangaController{service: service, cloudinaryService: cloudinaryService}
}

func (mangaController *MangaController) CreateManga(ctx *gin.Context) {
	var dto dtos.CreateMangaDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	manga, err := mangaController.service.CreateManga(context.Background(), &dto)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			ctx.JSON(http.StatusConflict, utils.ErrorResponse{
				Status:  http.StatusConflict,
				Error:   http.StatusText(http.StatusConflict),
				Message: "title already exists",
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

	resp := dtos.MangaResponseDTO{
		ID:          manga.ID.Hex(),
		Title:       manga.Title,
		Description: manga.Description,
		Author:      manga.Author,
		CoverImage:  manga.CoverImage,
		Genres:      manga.Genres,
		Status:      manga.Status,
		CreatedAt:   manga.CreatedAt,
		UpdatedAt:   manga.UpdatedAt,
	}

	ctx.JSON(201, utils.SuccessResponse[dtos.MangaResponseDTO]{
		Status:  201,
		Message: "manga created successfully",
		Data:    resp,
	})
}

func (mangaControler *MangaController) GetAllManga(ctx *gin.Context) {
	mangas, err := mangaControler.service.GetAllManga(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   http.StatusText(http.StatusInternalServerError),
			Message: "internal server error",
		})
		return
	}

	var mangaResponse []dtos.MangaResponseDTOWithGenre
	for _, result := range mangas {
		mangaResponse = append(mangaResponse, dtos.MangaResponseDTOWithGenre{
			ID:          result.ID.Hex(),
			Title:       result.Title,
			Description: result.Description,
			Author:      result.Author,
			CoverImage:  result.CoverImage,
			Genres:      result.Genres,
			Status:      result.Status,
			CreatedAt:   result.CreatedAt,
			UpdatedAt:   result.UpdatedAt,
		})
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse[[]dtos.MangaResponseDTOWithGenre]{
		Status:  http.StatusOK,
		Message: "mangas fetched successfully",
		Data:    mangaResponse,
	})
}

func (mangaController *MangaController) GetMangaById(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	manga, err := mangaController.service.GetMangaById(ctx, id)
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
			Message: "internal server error",
		})
		return
	}

	mangaResponse := dtos.MangaResponseDTOWithGenre{
		ID:          manga.ID.Hex(),
		Title:       manga.Title,
		Description: manga.Description,
		Author:      manga.Author,
		CoverImage:  manga.CoverImage,
		Genres:      manga.Genres,
		Status:      manga.Status,
		CreatedAt:   manga.CreatedAt,
		UpdatedAt:   manga.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse[dtos.MangaResponseDTOWithGenre]{
		Status:  http.StatusOK,
		Message: "manga found successfully",
		Data:    mangaResponse,
	})
}

func (mangaController *MangaController) FilterManga(ctx *gin.Context) {
	title := ctx.Query("title")
	genres := ctx.QueryArray("genres")
	status := ctx.Query("status")

	mangas, err := mangaController.service.FilterManga(ctx, title, genres, status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   http.StatusText(http.StatusInternalServerError),
			Message: "internal server error",
		})
		return
	}

	var mangaResponse []dtos.MangaResponseDTOWithGenre
	for _, result := range *mangas {
		mangaResponse = append(mangaResponse, dtos.MangaResponseDTOWithGenre{
			ID:          result.ID.Hex(),
			Title:       result.Title,
			Description: result.Description,
			Author:      result.Author,
			CoverImage:  result.CoverImage,
			Genres:      result.Genres,
			Status:      result.Status,
			CreatedAt:   result.CreatedAt,
			UpdatedAt:   result.UpdatedAt,
		})
	}

	if len(mangaResponse) == 0 {
		ctx.JSON(http.StatusOK, utils.SuccessResponse[[]dtos.MangaResponseDTOWithGenre]{
			Status:  http.StatusOK,
			Message: "no manga found matching filters",
			Data:    []dtos.MangaResponseDTOWithGenre{},
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse[[]dtos.MangaResponseDTOWithGenre]{
		Status:  http.StatusOK,
		Message: "mangas found successfully",
		Data:    mangaResponse,
	})
}
