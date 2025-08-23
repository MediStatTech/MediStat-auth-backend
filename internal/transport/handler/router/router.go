package router

import (
	"github.com/MediStatTech/MediStat-auth-backend/internal/services"
	"github.com/MediStatTech/MediStat-auth-backend/internal/transport"
	"github.com/MediStatTech/MediStat-auth-backend/internal/transport/handler"
)

func RegisterRoutes(server *transport.Server, services *services.Services) {
	// Register handlers
	handler.NewAuth(services).Register(server)

	// server.GET("/swagger/*", echoSwagger.WrapHandler)
	// server.GET("/health", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "OK")
	// })
}
