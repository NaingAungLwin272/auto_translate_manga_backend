package modules

import (
	"auto_translate_manga_backend/internal/services"
	"auto_translate_manga_backend/internal/utils"

	"go.uber.org/fx"
)

var ServiceModule = fx.Options(
	fx.Provide(
		utils.GetCloudinaryClient,
		services.NewUserService,
		services.NewGenreService,
		services.NewCloudinaryService,
		services.NewMangaService,
		services.NewFavoriteService,
		services.NewChapterService,
	),
)
