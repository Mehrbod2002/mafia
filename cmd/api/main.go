package main

import (
	"mafia/config"
	httpadapter "mafia/internal/adapters/http"
	"mafia/internal/adapters/postgres"
	"mafia/internal/adapters/redis"
	"mafia/internal/adapters/webrtc"
	"mafia/internal/core/services"
	"mafia/internal/ports"
	"mafia/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.Load()
	logger.Init(cfg.Logging.Level)

	db := postgres.New(cfg.Database.URL)
	cache := redis.New(cfg.Redis.Addr)
	sfu := webrtc.NewSFU()

	repos := ports.Repositories{
		User:      postgres.NewUserRepository(db),
		Room:      postgres.NewRoomRepository(db),
		Group:     postgres.NewGroupRepository(db),
		Wallet:    postgres.NewWalletRepository(db),
		Challenge: postgres.NewChallengeRepository(db),
		Role:      postgres.NewRoleRepository(db),
		Shop:      postgres.NewShopRepository(db),
		Rule:      postgres.NewRuleRepository(db),
		Scenario:  postgres.NewScenarioRepository(db),
	}

	services := services.NewServices(repos, cache, sfu)

	r := gin.Default()
	httpadapter.SetupRoutes(r, services, sfu)

	srv := &http.Server{Addr: ":" + cfg.Server.Port, Handler: r}
	go srv.ListenAndServe()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Info("Shutting down server...")
}
