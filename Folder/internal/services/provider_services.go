package services

import (
	"context"

	"github.com/go-kit/kit/log"
)

type ProviderService interface {
	FindAll(ctx context.Context) ([]entity.Provider, error)
	GetByID(ctx context.Context, id string) (entity.Provider, error)
	Create(ctx context.Context, Provider entity.Provider) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string) (entity.Provider, error)
}

type providerService struct {
	repository entity.ProviderRepository
	logger     log.Logger
}

func NewProviderService(repository entity.ProviderRepository, logger log.Logger) ProviderService {
	return &providerService{
		repository: repository,
		logger:     logger,
	}
}

func (s *providerService) FindAll(ctx context.Context) ([]entity.Provider, error) {
	provider, err := s.repository.FindAll(ctx)
	if err != nil {
		s.logger.Log("error:", err)
		return nil, err
	}
	s.logger.Log("findall:", "success")
	return provider, nil
}

func (s *providerService) GetByID(ctx context.Context, id string) (entity.Provider, error) {
	provider, err := s.repository.GetByID(ctx, id)
	if err != nil {
		s.logger.Log("error:", err)
		return provider, err
	}
	s.logger.Log("getbyid:", "success")
	return provider, nil
}

func (s *providerService) Create(ctx context.Context, provider entity.Provider) error {
	provider.Prepare()
	err := provider.Validate()
	if err != nil {
		s.logger.Log("error:", err)
		return err
	}

	if err := s.repository.Create(ctx, provider); err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("create:", "success")

	return nil
}

func (s *providerService) Delete(ctx context.Context, id string) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("delete:", "success")
	return nil
}
