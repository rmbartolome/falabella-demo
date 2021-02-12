package service

import (
	"context"
	"folder/internal/repository"

	"github.com/go-kit/kit/log"
)

type ProductService interface {
	FindAll(ctx context.Context) ([]repository.Product, error)
	GetByID(ctx context.Context, id string) (repository.Product, error)
	Create(ctx context.Context, Product repository.Product) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string) (repository.Product, error)
}

type productService struct {
	repository repository.ProductRepository
	logger     log.Logger
}

func NewProductService(repository repository.ProductRepository, logger log.Logger) ProductService {
	return &productService{
		repository: repository,
		logger:     logger,
	}
}

func (s *productService) FindAll(ctx context.Context) ([]repository.Product, error) {
	products, err := s.repository.FindAll(ctx)
	if err != nil {
		s.logger.Log("error:", err)
		return nil, err
	}
	s.logger.Log("findall:", "success")
	return products, nil
}

func (s *productService) GetByID(ctx context.Context, id string) (repository.Product, error) {
	products, err := s.repository.GetByID(ctx, id)
	if err != nil {
		s.logger.Log("error:", err)
		return products, err
	}
	s.logger.Log("getbyid:", "success")
	return products, nil
}

func (s *productService) Create(ctx context.Context, product repository.Product) error {
	product.Prepare()
	err := product.Validate()
	if err != nil {
		s.logger.Log("error:", err)
		return err
	}

	if err := s.repository.Create(ctx, product); err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("create:", "success")

	return nil
}

func (s *productService) Delete(ctx context.Context, id string) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("delete:", "success")
	return nil
}
