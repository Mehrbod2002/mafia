package main

import (
	"mafia/config"
	httpadapter "mafia/internal/adapters/http"
	"mafia/internal/adapters/postgres"
	"mafia/internal/adapters/webrtc"
	"mafia/internal/core/services"
	"mafia/internal/ports"
	cachepkg "mafia/pkg/cache"
	"mafia/pkg/events"
	"mafia/pkg/logger"
	"mafia/pkg/notifications"
	"mafia/pkg/payment"
	"mafia/pkg/queue"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"mafia/docs"
)

func main() {
	cfg := config.Load()
	logger.Init(cfg.Logging.Level)

	docs.SwaggerInfo.Host = "localhost:" + cfg.Server.Port
	docs.SwaggerInfo.BasePath = "/"

	db := postgres.New(cfg.Database.URL)
	inMemoryCache := cachepkg.NewInMemoryCache()
	taskQueue := queue.NewBackgroundQueue(128)
	eventBus := events.NewSimpleBus(taskQueue.Enqueue)
	notifier := notifications.NewLogSender()
	paymentProvider := payment.NewZarinpalProvider(cfg.Payment.Zarinpal)
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

	infra := ports.Infrastructure{
		Cache:         inMemoryCache,
		Queue:         taskQueue,
		Events:        eventBus,
		Notifications: notifier,
		Payments:      paymentProvider,
	}

	services := services.NewServices(repos, infra, sfu)

	r := gin.Default()
	httpadapter.SetupRoutes(r, services, sfu)

	srv := &http.Server{Addr: ":" + cfg.Server.Port, Handler: r}
	go srv.ListenAndServe()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Info("Shutting down server...")

	taskQueue.Close()
}
