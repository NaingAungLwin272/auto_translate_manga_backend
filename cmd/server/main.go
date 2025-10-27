package main

import (
	"auto_translate_manga_backend/internal/configs"
	"auto_translate_manga_backend/internal/modules"
	"log"

	"github.com/joho/godotenv"

	"go.uber.org/fx"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	fx.New(
		fx.Provide(
			configs.NewConfig,
		),
		modules.RepoModule,
		modules.MongoModule,
		modules.ServiceModule,
		modules.ControllerModule,
		modules.RouterModule,
		fx.Invoke(modules.RunServer),
	).Run()
}
