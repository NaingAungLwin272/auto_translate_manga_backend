package routes

import (
	"auto_translate_manga_backend/internal/controllers"
	"auto_translate_manga_backend/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	userController *controllers.UserController,
	genreController *controllers.GenreController,
) {
	// Global middlewares
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.LoggerMiddleware())

	// User routes
	users := r.Group("/users")
	{
		users.POST("/create_user", userController.CreateUser)
		users.GET("/get_users", userController.GetAllUsers)
	}

	// Genre routes
	genres := r.Group("/genres")
	{
		genres.POST("/create_genre", genreController.CreateGenre)
		genres.GET("/get_genres", genreController.GetAllGenres)
		genres.GET("/get_genre_by_id/:id", genreController.GetGenreById)
		genres.PATCH("/update_genre_by_id/:id", genreController.UpdateGenreById)
		genres.DELETE("/delete_genre_by_id/:id", genreController.DeleteGenreById)
	}
}
