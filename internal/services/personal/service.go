package personal

import (
	"github.com/MediStatTech/MediStat-auth-backend/internal/repository/personal"
	"github.com/MediStatTech/MediStat-auth-backend/internal/services/personal/add_personal"
	"github.com/MediStatTech/MediStat-auth-backend/internal/services/personal/get_personal"
	"github.com/MediStatTech/MediStat-auth-backend/internal/services/personal/list_personal"
	"github.com/MediStatTech/MediStat-auth-backend/internal/services/personal/personal_update_status"
)

type Service struct {
	AddPersonal          *add_personal.Facade
	ListPersonal         *list_personal.Facade
	UpdatePersonalStatus *personal_update_status.Facade
	GetPersonal          *get_personal.Facade
}

func NewService(perRepo *personal.Queries) *Service {
	return &Service{
		AddPersonal:          add_personal.New(perRepo),
		UpdatePersonalStatus: personal_update_status.New(perRepo),
		ListPersonal:         list_personal.New(perRepo),
		GetPersonal:          get_personal.New(perRepo),
	}
}
