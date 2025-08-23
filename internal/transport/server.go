package transport

import (
	"fmt"
	"net/http"

	"github.com/MediStatTech/MediStat-auth-backend/internal/config"
	"github.com/MediStatTech/MediStat-auth-backend/internal/transport/middleware"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Port string
	*echo.Echo
}

func NewServer(cfg *config.Config) *Server {
	server := echo.New()

	server.Use(middleware.RecoverMiddleware())
	server.Use(middleware.CORSMiddleware())

	server.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	return &Server{
		Port: cfg.HTTPPort,
		Echo: server,
	}
}

func (s *Server) Run() error {
	return s.Echo.Start(fmt.Sprintf(":%s", s.Port))
}
