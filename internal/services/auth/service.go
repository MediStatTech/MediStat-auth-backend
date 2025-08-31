package auth

import (
	"github.com/MediStatTech/MediStat-auth-backend/internal/repository/personal"
	"github.com/MediStatTech/MediStat-auth-backend/internal/services/auth/sign_in"
	jwt "github.com/MediStatTech/MediStat-jwt"
)

type Service struct {
	SignIn               *sign_in.Facade
}

func NewService(perRepo *personal.Queries, jwt *jwt.JWT) *Service {
	return &Service{
		SignIn:               sign_in.New(perRepo, jwt),
	}
}
