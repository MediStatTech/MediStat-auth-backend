package add_personal

import (
	"context"
	"database/sql"
	"time"

	"github.com/MediStatTech/MediStat-auth-backend/internal/domain/dto"
	"github.com/MediStatTech/MediStat-auth-backend/internal/domain/enums"
	"github.com/MediStatTech/MediStat-auth-backend/internal/repository/personal"
	"github.com/MediStatTech/MediStat-auth-backend/pkg/password"
	"github.com/MediStatTech/MediStat-auth-backend/pkg/uuid"
)

type service struct {
	ctx     context.Context
	req     *dto.AddPersonalRequest
	perRepo *personal.Queries

	passwordHash string
}

func (s *service) checkExists() error {
	ex, err := s.perRepo.ExistsPersonalByEmail(s.ctx, s.req.Email)
	if err != nil {
		return errFailedToFindUser.SetInternal(err)
	}

	if ex {
		return errUserAlreadyExists
	}

	return nil
}

func (s *service) hashPassword() error {
	hash, err := password.HashPassword(s.req.Password)
	if err != nil {
		return errFailedToHashPassword.SetInternal(err)
	}

	s.passwordHash = hash
	return nil
}

func (s *service) create() error {
	departure := enums.PersonalDepartureEnum(s.req.Departure)
	if !departure.IsValid() {
		return errInvalidDeparture
	}

	_, err := s.perRepo.CreatePersonal(s.ctx, personal.CreatePersonalParams{
		PersonalID:   uuid.NewUUID(),
		FirstName:    s.req.FirstName,
		LastName:     s.req.LastName,
		Email:        s.req.Email,
		Phone:        sql.NullString{String: *s.req.Phone, Valid: s.req.Phone != nil},
		PasswordHash: s.passwordHash,
		Status:       enums.PersonalStatusEnumActive.String(),
		Departure:    departure.String(),
		CreatedAt:    time.Now(),
		UpdatedAt:    sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return errFailedToCreatePersonal.SetInternal(err)
	}
	return nil
}

func (s *service) reply() (*dto.AddPersonalResponse, error) {
	return &dto.AddPersonalResponse{}, nil
}
