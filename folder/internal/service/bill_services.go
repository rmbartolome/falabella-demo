package service

import (
	"context"
	"folder/internal/repository"

	"github.com/go-kit/kit/log"
)

type BillService interface {
	FindAll(ctx context.Context) ([]repository.Bill, error)
	GetByID(ctx context.Context, id string) (repository.Bill, error)
	Create(ctx context.Context, Bill repository.Bill) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string) (repository.Bill, error)
}

type billService struct {
	repository repository.BillRepository
	logger     log.Logger
}

func NewBillService(repository repository.BillRepository, logger log.Logger) BillService {
	return &billService{
		repository: repository,
		logger:     logger,
	}
}

func (s *billService) FindAll(ctx context.Context) ([]repository.Bill, error) {
	bills, err := s.repository.FindAll(ctx)
	if err != nil {
		s.logger.Log("error:", err)
		return nil, err
	}
	s.logger.Log("findall:", "success")
	return bills, nil
}

func (s *billService) GetByID(ctx context.Context, id string) (repository.Bill, error) {
	bills, err := s.repository.GetByID(ctx, id)
	if err != nil {
		s.logger.Log("error:", err)
		return bills, err
	}
	s.logger.Log("getbyid:", "success")
	return bills, nil
}

func (s *billService) Create(ctx context.Context, bill repository.Bill) error {
	bill.Prepare()
	err := bill.Validate()
	if err != nil {
		s.logger.Log("error:", err)
		return err
	}

	if err := s.repository.Create(ctx, bill); err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("create:", "success")

	return nil
}

func (s *billService) Delete(ctx context.Context, id string) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("delete:", "success")
	return nil
}
