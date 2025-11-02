package modules

import (
	"auto_translate_manga_backend/internal/controllers"

	"go.uber.org/fx"
)

var ControllerModule = fx.Options(
	fx.Provide(
		controllers.NewUserController,
		controllers.NewGenreController,
		controllers.NewCloudinaryController,
		controllers.NewMangaController,
	),
)
