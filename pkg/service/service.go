package service

import "github.com/rostis232/traineeEVOFintech/pkg/repository"

type 

type Service struct {
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
