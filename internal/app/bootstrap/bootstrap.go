package bootstrap

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MediStatTech/MediStat-auth-backend/internal/config"
	"github.com/MediStatTech/MediStat-auth-backend/internal/repository"
	"github.com/MediStatTech/MediStat-auth-backend/internal/services"
	"github.com/MediStatTech/MediStat-auth-backend/internal/transport"
	"github.com/MediStatTech/MediStat-auth-backend/internal/transport/handler/router"
	"github.com/MediStatTech/MediStat-auth-backend/pkg/db"
	jwt "github.com/MediStatTech/MediStat-jwt"
	"github.com/MediStatTech/MediStat-log/logger"
)

func Run() {
	log := logger.New(os.Stdout)
	ctx, cancelFunc := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	defer cancelFunc()

	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	db, err := db.NewPostgresDB(cfg.PostgresDSN)
	if err != nil {
		panic(err)
	}

	jwt := jwt.New(cfg.JWTSecret, 72*time.Hour)

	rep := repository.New(db)

	serv := transport.NewServer(cfg)
	services := services.NewServices(rep, jwt)

	router.RegisterRoutes(serv, services, jwt)

	log.Info("MediStat Auth Backend is running", map[string]any{
		"port": cfg.HTTPPort,
	})

	go func() {
		if err := serv.Run(); err != nil {
			log.Error("server error", map[string]any{
				"error": err,
			})
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Info("MediStat Auth Backend is shutting down", map[string]any{})

	if err = serv.Shutdown(ctx); err != nil {
		log.Error("error with shutting down server", map[string]any{
			"error": err,
		})
	}
}
