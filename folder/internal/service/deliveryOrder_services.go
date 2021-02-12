package service

import (
	"context"
	"folder/internal/repository"

	"github.com/go-kit/kit/log"
)

type DeliveryOrderService interface {
	FindAll(ctx context.Context) ([]repository.DeliveryOrder, error)
	GetByID(ctx context.Context, id string) (repository.DeliveryOrder, error)
	Create(ctx context.Context, DeliveryOrder repository.DeliveryOrder) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string) (repository.DeliveryOrder, error)
}

type deliveryOrderService struct {
	repository repository.DeliveryOrderRepository
	logger     log.Logger
}

func NewDeliveryOrderService(repository repository.DeliveryOrderRepository, logger log.Logger) DeliveryOrderService {
	return &deliveryOrderService{
		repository: repository,
		logger:     logger,
	}
}

func (s *deliveryOrderService) FindAll(ctx context.Context) ([]repository.DeliveryOrder, error) {
	deliveryOrder, err := s.repository.FindAll(ctx)
	if err != nil {
		s.logger.Log("error:", err)
		return nil, err
	}
	s.logger.Log("findall:", "success")
	return deliveryOrder, nil
}

func (s *deliveryOrderService) GetByID(ctx context.Context, id string) (repository.DeliveryOrder, error) {
	deliveryOrder, err := s.repository.GetByID(ctx, id)
	if err != nil {
		s.logger.Log("error:", err)
		return deliveryOrder, err
	}
	s.logger.Log("getbyid:", "success")
	return deliveryOrder, nil
}

func (s *deliveryOrderService) Create(ctx context.Context, deliveryOrder repository.DeliveryOrder) error {
	deliveryOrder.Prepare()
	err := deliveryOrder.Validate()
	if err != nil {
		s.logger.Log("error:", err)
		return err
	}

	if err := s.repository.Create(ctx, deliveryOrder); err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("create:", "success")

	return nil
}

func (s *deliveryOrderService) Delete(ctx context.Context, id string) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("delete:", "success")
	return nil
}
