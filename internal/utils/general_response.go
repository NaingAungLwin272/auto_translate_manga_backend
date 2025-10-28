package utils

type SuccessResponse[T any] struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

type SuccessResponseForDeletingProcess[T any] struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
