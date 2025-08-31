package handler

import (
	"net/http"

	"github.com/MediStatTech/MediStat-auth-backend/internal/domain/dto"
	"github.com/MediStatTech/MediStat-auth-backend/internal/services"
	"github.com/MediStatTech/MediStat-auth-backend/internal/transport"
	"github.com/MediStatTech/MediStat-bind"
	"github.com/labstack/echo/v4"
)

type Auth struct {
	service *services.Services
}

func NewAuth(service *services.Services) *Auth {
	return &Auth{
		service: service,
	}
}

func (h *Auth) Register(serv *transport.Server) {
	group := serv.Group("/auth")

	group.POST("/sign-in", h.SignIn)
}

func (h *Auth) SignIn(c echo.Context) error {
	var (
		err error
		obj dto.SignInRequest
	)

	if err = bind.Validate(c, &obj); err != nil {
		return err
	}

	res, err := h.service.Auth.SignIn.Handle(c.Request().Context(), &obj)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
