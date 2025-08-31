package list_personal

import (
	"context"

	"github.com/MediStatTech/MediStat-auth-backend/internal/domain/dto"
	"github.com/MediStatTech/MediStat-auth-backend/internal/repository/personal"
)

type service struct {
	ctx     context.Context
	req     *dto.ListPersonalRequest
	perRepo *personal.Queries

	personal []personal.Personal
}

func (s *service) list() error {
	pers, err := s.perRepo.GetAllPersonal(s.ctx)
	if err != nil {
		return errFailedToListPersonal.SetInternal(err)
	}

	s.personal = pers
	return nil
}

func (s *service) reply() (*dto.ListPersonalResponse, error) {
	personals := make([]dto.Personal, len(s.personal))
	for i, p := range s.personal {
		personals[i] = dto.Personal{
			PersonalID: p.PersonalID.String(),
			FirstName:  p.FirstName,
			LastName:   p.LastName,
			Email:      p.Email,
			Phone:      &p.Phone.String,
			Status:     p.Status,
			Departure:  p.Departure,
			CreatedAt:  p.CreatedAt,
			UpdatedAt:  p.UpdatedAt.Time,
		}
	}
	return &dto.ListPersonalResponse{Personal: personals}, nil
}
