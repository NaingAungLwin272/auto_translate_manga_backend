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

type FavoriteController struct {
	service *services.FavoriteService
}

func NewFavoriteController(service *services.FavoriteService) *FavoriteController {
	return &FavoriteController{service: service}
}

func (favoriteController *FavoriteController) CreateFavoriteManga(ctx *gin.Context) {
	var dto dtos.CreateFavoritedDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   http.StatusText(http.StatusBadRequest),
			Message: "name must be require",
		})
		return
	}

	favorite, err := favoriteController.service.CreateFavoriteManga(context.Background(), &dto)
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

	favoriteResponse := dtos.FavoritedResponseDTO{
		ID:    favorite.ID,
		User:  favorite.User,
		Manga: favorite.Manga,
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse[dtos.FavoritedResponseDTO]{
		Status:  http.StatusOK,
		Message: "genre created successfully",
		Data:    favoriteResponse,
	})
}

func (favoriteController *FavoriteController) GetFavoriteMangaByUserId(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   http.StatusText(http.StatusBadRequest),
			Message: "id parameter is required",
		})
		return
	}

	favoriteManga, err := favoriteController.service.GetFavoriteMangaByUserId(ctx, id)
	if err != nil {
		if err.Error() == "no favorites found for this user" {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse{
				Status:  http.StatusNotFound,
				Error:   http.StatusText(http.StatusNotFound),
				Message: "no favorites found for this user",
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

	ctx.JSON(http.StatusOK, utils.SuccessResponse[*models.FavoriteWithUsersMangaDetail]{
		Status:  http.StatusOK,
		Message: "favorites found successfully",
		Data:    favoriteManga,
	})
}

func (favoriteController *FavoriteController) RemoveFavoriteMangaByUserId(ctx *gin.Context) {
	userId := ctx.Params.ByName("id")
	mangaId := ctx.Query("manga_id")

	if userId == "" {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   http.StatusText(http.StatusBadRequest),
			Message: "id parameter is required",
		})
		return
	}

	if mangaId == "" {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   http.StatusText(http.StatusBadRequest),
			Message: "manga query parameter is required",
		})
		return
	}

	mangas := strings.Split(mangaId, ",")
	result, err := favoriteController.service.RemoveFavoriteMangaByUserId(ctx, userId, mangas)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
		return
	}

	if strings.Contains(result, "remove unsuccessfully") {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   http.StatusText(http.StatusInternalServerError),
			Message: "favorite remove unsuccessfully",
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponseForDeletingProcess[any]{
		Status:  http.StatusOK,
		Message: "favorite removed successfully",
	})
}
