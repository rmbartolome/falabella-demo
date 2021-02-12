package service

import (
	"context"
	"folder/internal/repository"

	"github.com/go-kit/kit/log"
)

type ClientServices interface {
	FindAll(ctx context.Context) ([]repository.Client, error)
	GetByID(ctx context.Context, id string) (repository.Client, error)
	Create(ctx context.Context, Client repository.Client) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string) (repository.Client, error)
}

type clientService struct {
	repository repository.ClientRepository
	logger     log.Logger
}

func NewClientService(repository repository.ClientRepository, logger log.Logger) ClientServices {
	return &clientService{
		repository: repository,
		logger:     logger,
	}
}

func (s *clientService) FindAll(ctx context.Context) ([]repository.Client, error) {
	clients, err := s.repository.FindAll(ctx)
	if err != nil {
		s.logger.Log("error:", err)
		return nil, err
	}
	s.logger.Log("findall:", "success")
	return clients, nil
}

func (s *clientService) GetByID(ctx context.Context, id string) (repository.Client, error) {
	clients, err := s.repository.GetByID(ctx, id)
	if err != nil {
		s.logger.Log("error:", err)
		return clients, err
	}
	s.logger.Log("getbyid:", "success")
	return clients, nil
}

func (s *clientService) Create(ctx context.Context, client repository.Client) error {
	client.Prepare()
	err := client.Validate()
	if err != nil {
		s.logger.Log("error:", err)
		return err
	}

	if err := s.repository.Create(ctx, client); err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("create:", "success")

	return nil
}

func (s *clientService) Delete(ctx context.Context, id string) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("delete:", "success")
	return nil
}
