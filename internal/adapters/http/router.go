package http

import (
	"mafia/internal/ports"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, s ports.Services, _ ports.SFU) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", RegisterHandler(s.User))
		auth.POST("/verify", VerifyOTPHandler(s.User))
		auth.POST("/login", LoginHandler(s.User))
	}

	user := r.Group("/user").Use(AuthMiddleware(s.User))
	{
		user.GET("/profile", GetProfileHandler(s.User))
		user.PUT("/profile", UpdateProfileHandler(s.User))
		user.GET("/wallet", GetWalletHandler(s.Wallet))
		user.POST("/purchase", PurchaseHandler(s.Wallet))
	}
}
