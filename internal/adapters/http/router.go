package http

import (
	"mafia/internal/ports"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, s ports.Services, sfu ports.SFU) {
	// 100+ routes
	auth := r.Group("/auth")
	{ /* 20+ */ }
	user := r.Group("/user").Use(AuthMiddleware())
	{ /* 30+ */ }
	game := r.Group("/game").Use(AuthMiddleware())
	{ /* 25+ */ }
	// ... more
}
