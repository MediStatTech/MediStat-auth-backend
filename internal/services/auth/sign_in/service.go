package sign_in

import (
	"context"

	"github.com/MediStatTech/MediStat-auth-backend/internal/domain/dto"
	"github.com/MediStatTech/MediStat-auth-backend/internal/repository/personal"
	"github.com/MediStatTech/MediStat-jwt"
	"github.com/MediStatTech/MediStat-auth-backend/pkg/password"
	"github.com/MediStatTech/MediStat-auth-backend/internal/domain/enums"
)

type service struct {
	ctx     context.Context
	req     *dto.SignInRequest
	perRepo *personal.Queries
	jwt     *jwt.JWT

	personal personal.Personal
	token    string
}

func (s *service) findByEmail() error {
	per, err := s.perRepo.GetPersonalByEmail(s.ctx, s.req.Email)
	if err != nil {
		return errFailedFindUser.SetInternal(err)
	}

	if per.Status == enums.PersonalStatusEnumInactive.String() {
		return errInactivePersonal
	}

	s.personal = per
	return nil
}

func (s *service) checkPassword() error {
	if !password.CheckPasswordHash(s.req.Password, s.personal.PasswordHash) {
		return errInvalidPassword
	}

	return nil
}

func (s *service) generateToken() error {
	var err error

	if s.token, err  = s.jwt.Generate(s.personal.PersonalID.String()); err != nil {
		return errFailedGenerateToken.SetInternal(err)
	}

	return nil
}

func (s *service) reply() (*dto.SignInResponse, error) {
	return &dto.SignInResponse{Token: s.token}, nil
}
