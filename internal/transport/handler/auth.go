package handler

import (
	"fmt"

	"github.com/MediStatTech/MediStat-auth-backend/internal/services"
	"github.com/MediStatTech/MediStat-auth-backend/internal/transport"
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
	fmt.Println("auth group", group)
}
