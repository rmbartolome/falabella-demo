package service

import (
	"context"
	"folder/internal/repository"

	"github.com/go-kit/kit/log"
)

type StockServices interface {
	FindAll(ctx context.Context) ([]repository.Stock, error)
	GetByID(ctx context.Context, id string) (repository.Stock, error)
	Create(ctx context.Context, Stock repository.Stock) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string) (repository.Stock, error)
}

type stockService struct {
	repository repository.StockRepository
	logger     log.Logger
}

func NewStockService(repository repository.StockRepository, logger log.Logger) StockServices {
	return &stockService{
		repository: repository,
		logger:     logger,
	}
}

func (s *stockService) FindAll(ctx context.Context) ([]repository.Stock, error) {
	stocks, err := s.repository.FindAll(ctx)
	if err != nil {
		s.logger.Log("error:", err)
		return nil, err
	}
	s.logger.Log("findall:", "success")
	return stocks, nil
}

func (s *stockService) GetByID(ctx context.Context, id string) (repository.Stock, error) {
	stocks, err := s.repository.GetByID(ctx, id)
	if err != nil {
		s.logger.Log("error:", err)
		return stocks, err
	}
	s.logger.Log("getbyid:", "success")
	return stocks, nil
}

func (s *stockService) Create(ctx context.Context, stock repository.Stock) error {
	stock.Prepare()
	err := stock.Validate()
	if err != nil {
		s.logger.Log("error:", err)
		return err
	}

	if err := s.repository.Create(ctx, stock); err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("create:", "success")

	return nil
}

func (s *stockService) Delete(ctx context.Context, id string) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("delete:", "success")
	return nil
}
