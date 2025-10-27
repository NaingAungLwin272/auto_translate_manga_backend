package modules

import (
	"context"
	"log"
	"time"

	"auto_translate_manga_backend/internal/configs"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"go.uber.org/fx"
)

func NewMongoClient(cfg *configs.Config, lc fx.Lifecycle) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(cfg.MongoURI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			if err := client.Ping(ctx, readpref.Primary()); err != nil {
				return err
			}
			log.Println("✅ Connected to MongoDB")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			if err := client.Disconnect(ctx); err != nil {
				return err
			}
			log.Println("🛑 Disconnected from MongoDB")
			return nil
		},
	})

	return client, nil
}

var MongoModule = fx.Options(
	fx.Provide(NewMongoClient),
)
