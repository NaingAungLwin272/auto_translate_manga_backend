package modules

import (
	"auto_translate_manga_backend/internal/controllers"
	"auto_translate_manga_backend/internal/routes"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewRouter(
	userController *controllers.UserController,
	genreController *controllers.GenreController,
	mangaController *controllers.MangaController,
	cloudinaryController *controllers.CloudinaryController,
) *gin.Engine {
	r := gin.Default()
	routes.SetupRoutes(r, userController, genreController, mangaController, cloudinaryController)
	return r
}

var RouterModule = fx.Options(
	fx.Provide(NewRouter),
)
