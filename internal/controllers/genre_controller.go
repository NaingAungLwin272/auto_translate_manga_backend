package controllers

import (
	dtos "auto_translate_manga_backend/internal/dto"
	"auto_translate_manga_backend/internal/models"
	"auto_translate_manga_backend/internal/services"
	"auto_translate_manga_backend/internal/utils"
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type GenreController struct {
	service *services.GenreService
}

func NewGenreController(service *services.GenreService) *GenreController {
	return &GenreController{service: service}
}

func (genreController *GenreController) CreateGenre(ctx *gin.Context) {
	var dto dtos.CreateGenreDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   http.StatusText(http.StatusBadRequest),
			Message: "name must be require",
		})
		return
	}

	genre, err := genreController.service.CreateGenre(context.Background(), &dto)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			ctx.JSON(http.StatusConflict, utils.ErrorResponse{
				Status:  http.StatusConflict,
				Error:   http.StatusText(http.StatusConflict),
				Message: "name already exists",
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

	genreResponse := dtos.GenreResponseDTO{
		ID:          genre.ID.Hex(),
		Name:        genre.Name,
		Description: genre.Description,
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse[dtos.GenreResponseDTO]{
		Status:  http.StatusOK,
		Message: "genre created successfully",
		Data:    genreResponse,
	})
}

func (genreController *GenreController) GetAllGenres(ctx *gin.Context) {
	genres, err := genreController.service.GetAllGenres(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   http.StatusText(http.StatusInternalServerError),
			Message: "internal server error",
		})
		return
	}

	var genreResponses []dtos.GenreResponseDTO
	for _, result := range genres {
		genreResponses = append(genreResponses, dtos.GenreResponseDTO{
			ID:          result.ID.Hex(),
			Name:        result.Name,
			Description: result.Description,
		})
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse[[]dtos.GenreResponseDTO]{
		Status:  http.StatusOK,
		Message: "genres fetched successfully",
		Data:    genreResponses,
	})
}

func (genreController *GenreController) GetGenreById(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   http.StatusText(http.StatusBadRequest),
			Message: "id prarmeter must be require",
		})
	}

	genreData, err := genreController.service.GetGenreById(ctx, id)
	if err != nil {
		if err.Error() == "genre not found" {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
				Status:  http.StatusNotFound,
				Error:   http.StatusText(http.StatusNotFound),
				Message: "genre not found",
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

	genreResponse := dtos.GenreResponseDTO{
		ID:          genreData.ID.Hex(),
		Name:        genreData.Name,
		Description: genreData.Description,
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse[dtos.GenreResponseDTO]{
		Status:  http.StatusOK,
		Message: "genre found successfully",
		Data:    genreResponse,
	})

}

func (genreController *GenreController) UpdateGenreById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   http.StatusText(http.StatusBadRequest),
			Message: "id parameter is required",
		})
		return
	}

	var dto dtos.UpdateGenreDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   http.StatusText(http.StatusBadRequest),
			Message: "invalid request body",
		})
		return
	}

	genre, err := genreController.service.UpdateGenreById(context.Background(), id, &dto)
	if err != nil {
		if err.Error() == "genre not found" {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
				Status:  http.StatusNotFound,
				Error:   http.StatusText(http.StatusNotFound),
				Message: "genre not found",
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

	ctx.JSON(http.StatusOK, utils.SuccessResponse[models.Genre]{
		Status:  http.StatusOK,
		Message: "genre updated successfully",
		Data:    *genre,
	})
}

func (genreController *GenreController) DeleteGenreById(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   http.StatusText(http.StatusBadRequest),
			Message: "id parameter is required",
		})
		return
	}

	_, err := genreController.service.DeleteGenreById(context.Background(), id)
	if err != nil {
		if err.Error() == "genre not found" {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
				Status:  http.StatusNotFound,
				Error:   http.StatusText(http.StatusNotFound),
				Message: "genre not found",
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

	ctx.JSON(http.StatusOK, utils.SuccessResponseForDeletingProcess[models.Genre]{
		Status:  http.StatusOK,
		Message: "genre deleted successfully",
	})

}
