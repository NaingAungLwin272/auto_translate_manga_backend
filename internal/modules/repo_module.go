package modules

import (
	"auto_translate_manga_backend/internal/repositories"

	"go.uber.org/fx"
)

var RepoModule = fx.Options(
	fx.Provide(
		repositories.NewUserRepo,
		repositories.NewGenreRepo,
		repositories.NewMangaRepo,
	),
)
