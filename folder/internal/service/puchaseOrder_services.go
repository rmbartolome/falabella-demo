package service

import (
	"context"
	"folder/internal/repository"

	"github.com/go-kit/kit/log"
)

type PuchaseOrderService interface {
	FindAll(ctx context.Context) ([]repository.PuchaseOrder, error)
	GetByID(ctx context.Context, id string) (repository.PuchaseOrder, error)
	Create(ctx context.Context, PuchaseOrder repository.PuchaseOrder) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string) (repository.PuchaseOrder, error)
}

type puchaseOrderService struct {
	repository repository.PuchaseOrderRepository
	logger     log.Logger
}

func NewPuchaseOrderService(repository repository.PuchaseOrderRepository, logger log.Logger) PuchaseOrderService {
	return &puchaseOrderService{
		repository: repository,
		logger:     logger,
	}
}

func (s *puchaseOrderService) FindAll(ctx context.Context) ([]repository.PuchaseOrder, error) {
	puchaseOrder, err := s.repository.FindAll(ctx)
	if err != nil {
		s.logger.Log("error:", err)
		return nil, err
	}
	s.logger.Log("findall:", "success")
	return puchaseOrder, nil
}

func (s *puchaseOrderService) GetByID(ctx context.Context, id string) (repository.PuchaseOrder, error) {
	puchaseOrder, err := s.repository.GetByID(ctx, id)
	if err != nil {
		s.logger.Log("error:", err)
		return puchaseOrder, err
	}
	s.logger.Log("getbyid:", "success")
	return puchaseOrder, nil
}

func (s *puchaseOrderService) Create(ctx context.Context, puchaseOrder repository.PuchaseOrder) error {
	puchaseOrder.Prepare()
	err := puchaseOrder.Validate()
	if err != nil {
		s.logger.Log("error:", err)
		return err
	}

	if err := s.repository.Create(ctx, puchaseOrder); err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("create:", "success")

	return nil
}

func (s *puchaseOrderService) Delete(ctx context.Context, id string) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("delete:", "success")
	return nil
}
