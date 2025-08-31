package handler

import (
	"net/http"

	"github.com/MediStatTech/MediStat-auth-backend/internal/domain/dto"
	"github.com/MediStatTech/MediStat-auth-backend/internal/services"
	"github.com/MediStatTech/MediStat-auth-backend/internal/transport"
	"github.com/MediStatTech/MediStat-auth-backend/internal/transport/middleware"
	bind "github.com/MediStatTech/MediStat-bind"
	jwt "github.com/MediStatTech/MediStat-jwt"
	"github.com/labstack/echo/v4"
)

type Personal struct {
	service *services.Services
	jwt     *jwt.JWT
}

func NewPersonal(service *services.Services, jwt *jwt.JWT) *Personal {
	return &Personal{
		service: service,
		jwt:     jwt,
	}
}

func (h *Personal) Register(serv *transport.Server) {
	group := serv.Group("/personal", middleware.JWT(*h.jwt))

	group.POST("/add", h.AddPersonal)
	group.PUT("/status", h.UpdatePersonalStatus)
	group.GET("/you", h.GetPersonal)
	group.GET("", h.ListPersonal)
}

func (h *Personal) AddPersonal(c echo.Context) error {
	var (
		err error
		obj dto.AddPersonalRequest
	)

	if err = bind.Validate(c, &obj); err != nil {
		return err
	}

	res, err := h.service.Personal.AddPersonal.Handle(c.Request().Context(), &obj)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *Personal) UpdatePersonalStatus(c echo.Context) error {
	var (
		err error
		obj dto.UpdatePersonalStatusRequest
	)

	if err = bind.Validate(c, &obj); err != nil {
		return err
	}

	res, err := h.service.Personal.UpdatePersonalStatus.Handle(c.Request().Context(), &obj)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Personal) GetPersonal(c echo.Context) error {
	var (
		err error
		obj dto.GetPersonalRequest
	)
	
	if err = bind.Validate(c, &obj); err != nil {
		return err
	}

	res, err := h.service.Personal.GetPersonal.Handle(c.Request().Context(), &obj)
	if err != nil {
		return err
	}
	
	return c.JSON(http.StatusOK, res)
}

func (h *Personal) ListPersonal(c echo.Context) error {
	var (
		err error
		obj dto.ListPersonalRequest
	)
	
	if err = bind.Validate(c, &obj); err != nil {
		return err
	}

	res, err := h.service.Personal.ListPersonal.Handle(c.Request().Context(), &obj)
	if err != nil {
		return err
	}
	
	return c.JSON(http.StatusOK, res)
}