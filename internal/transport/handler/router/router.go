package router

import (
	"net/http"

	"github.com/MediStatTech/MediStat-auth-backend/internal/services"
	"github.com/MediStatTech/MediStat-auth-backend/internal/transport"
	"github.com/MediStatTech/MediStat-auth-backend/internal/transport/handler"
	jwt "github.com/MediStatTech/MediStat-jwt"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RegisterRoutes(server *transport.Server, services *services.Services, jwt *jwt.JWT) {
	// Register handlers
	handler.NewAuth(services).Register(server)
	handler.NewPersonal(services, jwt).Register(server)

	server.GET("/swagger/*", echoSwagger.WrapHandler)
	server.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
}
