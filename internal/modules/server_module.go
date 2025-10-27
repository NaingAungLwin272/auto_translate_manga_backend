package modules

import (
	"auto_translate_manga_backend/internal/configs"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func RunServer(lc fx.Lifecycle, cfg *configs.Config, r *gin.Engine) {
	server := &http.Server{
		Addr:    cfg.Port,
		Handler: r,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Printf("🚀 Server running on %s\n", cfg.Port)
			go func() {
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("listen: %s\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("🛑 Shutting down server...")
			return server.Shutdown(ctx)
		},
	})
}
