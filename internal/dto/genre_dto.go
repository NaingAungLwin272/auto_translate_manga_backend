package dtos

import "time"

type CreateGenreDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description,omitempty"`
}

type UpdateGenreDTO struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type GenreResponseDTO struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateGenreResponseDTO struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}
