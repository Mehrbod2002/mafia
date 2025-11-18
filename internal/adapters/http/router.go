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
		user.GET("/dashboard", DashboardHandler(s.User))
		user.GET("/wallet", GetWalletHandler(s.Wallet))
		user.POST("/purchase", PurchaseHandler(s.Wallet))
	}

	shop := r.Group("/shop").Use(AuthMiddleware(s.User))
	{
		shop.GET("/items", ListShopItemsHandler(s.Shop))
		shop.POST("/purchase", PurchaseItemHandler(s.Shop))
	}

	game := r.Group("/game").Use(AuthMiddleware(s.User))
	{
		game.POST("/rooms", CreateRoomHandler(s.Game))
		game.GET("/rooms", ListRoomsHandler(s.Game))
		game.POST("/rooms/:id/join", JoinRoomHandler(s.Game))
		game.POST("/rooms/:id/leave", LeaveRoomHandler(s.Game))
		game.POST("/rooms/:id/start", StartGameHandler(s.Game))
		game.POST("/rooms/:id/phase", AdvancePhaseHandler(s.Game))
		game.POST("/rooms/:id/vote", VoteHandler(s.Game))
		game.POST("/rooms/:id/ability", AbilityHandler(s.Game))
	}

	admin := r.Group("/admin").Use(AuthMiddleware(s.User), AdminMiddleware(s.User))
	{
		AdminRoleRoutes(admin, s.Admin)
		AdminRuleRoutes(admin, s.Admin)
		AdminScenarioRoutes(admin, s.Admin)
		admin.POST("/shop/items", CreateShopItemHandler(s.Shop))
		admin.PUT("/shop/items/:id", UpdateShopItemHandler(s.Shop))
		admin.DELETE("/shop/items/:id", DeleteShopItemHandler(s.Shop))
	}
}
