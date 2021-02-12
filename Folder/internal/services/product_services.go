package services

import (
	"context"

	"github.com/go-kit/kit/log"
)

type ProductService interface {
	FindAll(ctx context.Context) ([]entity.Product, error)
	GetByID(ctx context.Context, id string) (entity.Product, error)
	Create(ctx context.Context, Product entity.Product) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string) (entity.Product, error)
}

type productService struct {
	repository entity.ProductRepository
	logger     log.Logger
}

func NewProductService(repository entity.ProductRepository, logger log.Logger) ProductService {
	return &productService{
		repository: repository,
		logger:     logger,
	}
}

func (s *productService) FindAll(ctx context.Context) ([]entity.Product, error) {
	products, err := s.repository.FindAll(ctx)
	if err != nil {
		s.logger.Log("error:", err)
		return nil, err
	}
	s.logger.Log("findall:", "success")
	return products, nil
}

func (s *productService) GetByID(ctx context.Context, id string) (entity.Product, error) {
	products, err := s.repository.GetByID(ctx, id)
	if err != nil {
		s.logger.Log("error:", err)
		return products, err
	}
	s.logger.Log("getbyid:", "success")
	return products, nil
}

func (s *productService) Create(ctx context.Context, product entity.Product) error {
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
