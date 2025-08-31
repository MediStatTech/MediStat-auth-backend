package personal_update_status

import (
	"context"
	"github.com/MediStatTech/MediStat-auth-backend/internal/domain/dto"
	"github.com/MediStatTech/MediStat-auth-backend/internal/repository/personal"
)

type Facade struct {
	perRepo *personal.Queries
} 

func New(perRepo *personal.Queries) *Facade {
	return &Facade{
		perRepo: perRepo,
	}
}

func (f *Facade) Handle(ctx context.Context, req *dto.UpdatePersonalStatusRequest) (*dto.UpdatePersonalStatusResponse, error) {
	serv := service{
		ctx:     ctx,
		req:     req,
		perRepo: f.perRepo,
	}

	if err := serv.updateStatus(); err != nil {
		return nil, err
	}
	
	return serv.reply()
}