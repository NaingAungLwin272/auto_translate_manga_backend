package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware logs request info
func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next() // process request
		latency := time.Since(start)
		status := ctx.Writer.Status()
		log.Printf("[%d] %s %s in %v", status, ctx.Request.Method, ctx.Request.URL.Path, latency)
	}
}
