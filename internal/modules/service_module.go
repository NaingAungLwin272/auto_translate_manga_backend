package modules

import (
	"auto_translate_manga_backend/internal/services"

	"go.uber.org/fx"
)

var ServiceModule = fx.Options(
	fx.Provide(
		services.NewUserService,
		services.NewGenreService,
	),
)
