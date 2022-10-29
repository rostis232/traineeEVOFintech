package service

import (
	"github.com/rostis232/traineeEVOFintech"
	"github.com/rostis232/traineeEVOFintech/pkg/repository"
)

type TransactionService struct {
	repo repository.Transaction
}

func NewTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) InsertToDB(transactions []traineeEVOFintech.Transaction) error {
	return s.repo.InsertToDB(transactions)
}

func (s *TransactionService) GetJSON(m map[string]string) ([]traineeEVOFintech.Transaction, error) {
	return s.repo.GetJSON(m)
}
