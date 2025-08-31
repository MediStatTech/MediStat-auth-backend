package services

import (
	"github.com/MediStatTech/MediStat-auth-backend/internal/repository"
	"github.com/MediStatTech/MediStat-auth-backend/internal/services/auth"
	"github.com/MediStatTech/MediStat-auth-backend/internal/services/personal"
	jwt "github.com/MediStatTech/MediStat-jwt"
)

type Services struct {
	Auth     *auth.Service
	Personal *personal.Service
}

func NewServices(rep *repository.Repository, jwt *jwt.JWT) *Services {
	return &Services{
		Auth:     auth.NewService(rep.Personal, jwt),
		Personal: personal.NewService(rep.Personal),
	}
}
