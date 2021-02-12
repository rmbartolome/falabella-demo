package service

import (
	"context"
	"folder/internal/repository"

	"github.com/go-kit/kit/log"
)

type TypeProdServices interface {
	FindAll(ctx context.Context) ([]repository.TypeProd, error)
	GetByID(ctx context.Context, id string) (repository.TypeProd, error)
	Create(ctx context.Context, TypeProd repository.TypeProd) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string) (repository.TypeProd, error)
}

type typeProdService struct {
	repository repository.TypeProdRepository
	logger     log.Logger
}

func NewTypeProdService(repository repository.TypeProdRepository, logger log.Logger) TypeProdServices {
	return &typeProdService{
		repository: repository,
		logger:     logger,
	}
}

func (s *typeProdService) FindAll(ctx context.Context) ([]repository.TypeProd, error) {
	typeProds, err := s.repository.FindAll(ctx)
	if err != nil {
		s.logger.Log("error:", err)
		return nil, err
	}
	s.logger.Log("findall:", "success")
	return typeProds, nil
}

func (s *typeProdService) GetByID(ctx context.Context, id string) (repository.TypeProd, error) {
	typeProds, err := s.repository.GetByID(ctx, id)
	if err != nil {
		s.logger.Log("error:", err)
		return typeProds, err
	}
	s.logger.Log("getbyid:", "success")
	return typeProds, nil
}

func (s *typeProdService) Create(ctx context.Context, typeProd repository.TypeProd) error {
	typeProd.Prepare()
	err := typeProd.Validate()
	if err != nil {
		s.logger.Log("error:", err)
		return err
	}

	if err := s.repository.Create(ctx, typeProd); err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("create:", "success")

	return nil
}

func (s *typeProdService) Delete(ctx context.Context, id string) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("delete:", "success")
	return nil
}
