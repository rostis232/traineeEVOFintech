package service

import (
	"github.com/rostis232/traineeEVOFintech"
	"github.com/rostis232/traineeEVOFintech/pkg/repository"
)

type Transaction interface {
	InsertToDB(transactions []traineeEVOFintech.Transaction) error
	GetJSON(m map[string]string) ([]traineeEVOFintech.TransactionT, error)
}

type Service struct {
	Transaction
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Transaction: NewTransactionService(repos.Transaction),
	}
}
