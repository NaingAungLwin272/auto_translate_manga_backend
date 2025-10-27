package services

import (
	dtos "auto_translate_manga_backend/internal/dto"
	"auto_translate_manga_backend/internal/models"
	"auto_translate_manga_backend/internal/repositories"
	"context"
)

type UserService struct {
	repo *repositories.UserRepo
}

func NewUserService(repo *repositories.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (userService *UserService) CreateUser(ctx context.Context, dto *dtos.CreateUserDTO) (*models.User, error) {
	user := &models.User{
		Username: dto.Username,
		Email:    dto.Email,
		Age:      dto.Age,
	}
	return userService.repo.CreateUser(ctx, user)
}

func (userService *UserService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return userService.repo.GetAllUsers(ctx)
}
