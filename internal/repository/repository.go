package repository

import (
	"database/sql"

	"github.com/MediStatTech/MediStat-auth-backend/internal/repository/personal"
	"github.com/MediStatTech/MediStat-auth-backend/internal/repository/staff_roles"
)

type Repository struct {
	Personal   *personal.Queries
	StaffRoles *staff_roles.Queries
}

func New(db *sql.DB) *Repository {
	return &Repository{
		Personal:   personal.New(db),
		StaffRoles: staff_roles.New(db),
	}
}

func NewWithTx(tx *sql.Tx) *Repository {
	return &Repository{
		Personal:   personal.New(tx),
		StaffRoles: staff_roles.New(tx),
	}
}
