package routes

import (
	"auto_translate_manga_backend/internal/controllers"
	"auto_translate_manga_backend/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	userController *controllers.UserController,
) {
	// Global middlewares
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.LoggerMiddleware())

	// User routes
	users := r.Group("/users")
	{
		users.POST("/create_user", userController.CreateUser)
	}
}
