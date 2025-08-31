package personal_update_status

import (
	"context"
	"database/sql"
	"time"

	"github.com/MediStatTech/MediStat-auth-backend/internal/domain/dto"
	"github.com/MediStatTech/MediStat-auth-backend/internal/domain/enums"
	"github.com/MediStatTech/MediStat-auth-backend/internal/repository/personal"
	"github.com/MediStatTech/MediStat-auth-backend/pkg/uuid"
)

type service struct {
	ctx     context.Context
	req     *dto.UpdatePersonalStatusRequest
	perRepo *personal.Queries

	personal personal.Personal
}

func (s *service) updateStatus() error {
	var err error

	departure := enums.PersonalStatusEnum(s.req.Status)
	if !departure.IsValid() {
		return errInvalidDeparture
	}

	if s.personal, err = s.perRepo.UpdatePersonalStatus(s.ctx, personal.UpdatePersonalStatusParams{
		PersonalID: uuid.MustParse(s.req.PersonalID),
		Status:     departure.String(),
		UpdatedAt:  sql.NullTime{Time: time.Now().UTC(), Valid: true},
	}); err != nil {
		return errFailedToUpdatePersonalStatus.SetInternal(err)
	}

	return nil
}

func (s *service) reply() (*dto.UpdatePersonalStatusResponse, error) {
	return &dto.UpdatePersonalStatusResponse{Personal: dto.Personal{
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
