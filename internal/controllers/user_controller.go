package controllers

import (
	dtos "auto_translate_manga_backend/internal/dto"
	"auto_translate_manga_backend/internal/services"
	"auto_translate_manga_backend/internal/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var dto dtos.CreateUserDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := c.service.CreateUser(context.Background(), &dto)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	resp := dtos.UserResponseDTO{
		ID:       user.ID.Hex(),
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}

	ctx.JSON(201, utils.SuccessResponse[dtos.UserResponseDTO]{
		Status:  201,
		Message: "user created successfully",
		Data:    resp,
	})
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.service.GetAllUsers(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   http.StatusText(http.StatusInternalServerError),
			Message: "internal server error",
		})
		return
	}

	var resp []dtos.UserResponseDTO
	for _, u := range users {
		resp = append(resp, dtos.UserResponseDTO{
			ID:       u.ID.Hex(),
			Username: u.Username,
			Email:    u.Email,
		})
	}

	ctx.JSON(200, utils.SuccessResponse[[]dtos.UserResponseDTO]{
		Status:  200,
		Message: "users fetched successfully",
		Data:    resp,
	})
}
