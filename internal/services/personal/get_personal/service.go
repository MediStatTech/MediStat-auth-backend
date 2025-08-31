package get_personal

import (
	"context"

	"github.com/MediStatTech/MediStat-auth-backend/internal/domain/dto"
	"github.com/MediStatTech/MediStat-auth-backend/internal/domain/enums"
	"github.com/MediStatTech/MediStat-auth-backend/internal/repository/personal"
	"github.com/MediStatTech/MediStat-auth-backend/pkg/uuid"
)

type service struct {
	ctx     context.Context
	req     *dto.GetPersonalRequest
	perRepo *personal.Queries

	personal personal.Personal
}

func (s *service) findByID() error {
	var err error

	if s.personal, err = s.perRepo.GetPersonalByID(s.ctx, uuid.MustParse(s.req.PersonalID)); err != nil {
		return errFailedToFindUser.SetInternal(err)
	}

	if s.personal.Status == enums.PersonalStatusEnumInactive.String() {
		return errInactivePersonal
	}

	return nil
}

func (s *service) reply() (*dto.GetPersonalResponse, error) {
	return &dto.GetPersonalResponse{Personal: dto.Personal{
		PersonalID: s.personal.PersonalID.String(),
		FirstName:  s.personal.FirstName,
		LastName:   s.personal.LastName,
		Email:      s.personal.Email,
		Phone:      &s.personal.Phone.String,
		Status:     s.personal.Status,
		Departure:  s.personal.Departure,
		CreatedAt:  s.personal.CreatedAt,
		UpdatedAt:  s.personal.UpdatedAt.Time,
	}}, nil
}
