package service

import "github.com/sergkim13/short-link-backend-go/pkg/repository"

type Link interface {

}

type Service struct {
	Link
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}