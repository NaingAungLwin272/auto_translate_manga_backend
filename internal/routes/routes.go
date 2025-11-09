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
	mangaController *controllers.MangaController,
	favoriteController *controllers.FavoriteController,
	chapterController *controllers.ChapterController,
	cloudinaryController *controllers.CloudinaryController,
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

	// Manga routes
	mangas := r.Group("/mangas")
	{
		mangas.POST("/create_manga", mangaController.CreateManga)
		mangas.GET("/get_all_manga", mangaController.GetAllManga)
		mangas.GET("/get_manga_by_id/:id", mangaController.GetMangaById)
		mangas.GET("/filter_manga", mangaController.FilterManga)
		mangas.DELETE("/delete_manga/:id", mangaController.DeleteManga)
	}

	// Favorite manga by user routes

	favorites := r.Group("/favorites")
	{
		favorites.POST("/create_favorite_manga", favoriteController.CreateFavoriteManga)
		favorites.GET("/get_favorite_manga_by_user_id/:id", favoriteController.GetFavoriteMangaByUserId)
		favorites.DELETE("/remove_favorite_manga/:id", favoriteController.RemoveFavoriteMangaByUserId)
	}

	chapters := r.Group("/chapters")
	{
		chapters.POST("/create_chapter", chapterController.CreateChapter)
	}

	// Cloudinary cover_image upload for manga
	cover_image := r.Group("/file_upload")
	{
		cover_image.POST("/cover_image_upload", cloudinaryController.UploadCoverImageToCloudinary)
	}
}
