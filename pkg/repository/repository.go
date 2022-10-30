package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rostis232/traineeEVOFintech"
)

type Transaction interface {
	InsertToDB(transactions []traineeEVOFintech.Transaction) error
	GetJSON(m map[string]string) ([]traineeEVOFintech.Transaction, error)
	GetCSV(m map[string]string) ([]traineeEVOFintech.Transaction, error)
}

type Repository struct {
	Transaction
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Transaction: NewTransactionPostgres(db),
	}
}
