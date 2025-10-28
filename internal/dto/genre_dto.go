package dtos

type CreateGenreDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description,omitempty"`
}

type UpdateGenreDTO struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type GenreResponseDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateGenreResponseDTO struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}
