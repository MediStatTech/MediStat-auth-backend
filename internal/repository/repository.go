package repository

import (
	"database/sql"
	"github.com/MediStatTech/MediStat-auth-backend/internal/repository/personal"
)

type Repository struct {
	Personal *personal.Queries
}

func New(db *sql.DB) *Repository {
	return &Repository{
		Personal: personal.New(db),
	}
}

func NewWithTx(tx *sql.Tx) *Repository {
	return &Repository{
		Personal: personal.New(tx),
	}
}
