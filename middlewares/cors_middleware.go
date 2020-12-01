package middlewares

import (
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"time"
)

func CorsMiddleware() gin.HandlerFunc {
	return cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	})
}
