package services

import (
	"context"

	"github.com/go-kit/kit/log"
)

type ClientServices interface {
	FindAll(ctx context.Context) ([]entity.Client, error)
	GetByID(ctx context.Context, id string) (entity.Bill, error)
	Create(ctx context.Context, Bill entity.Bill) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string) (entity.Bill, error)
}

type billService struct {
	repository entity.BillRepository
	logger     log.Logger
}

func NewClientService(repository entity.BillRepository, logger log.Logger) BillService {
	return &billService{
		repository: repository,
		logger:     logger,
	}
}

func (s *billService) FindAll(ctx context.Context) ([]entity.Bill, error) {
	bills, err := s.repository.FindAll(ctx)
	if err != nil {
		s.logger.Log("error:", err)
		return nil, err
	}
	s.logger.Log("findall:", "success")
	return bills, nil
}

func (s *billService) GetByID(ctx context.Context, id string) (entity.Bill, error) {
	bills, err := s.repository.GetByID(ctx, id)
	if err != nil {
		s.logger.Log("error:", err)
		return bills, err
	}
	s.logger.Log("getbyid:", "success")
	return bills, nil
}

func (s *billService) Create(ctx context.Context, bill entity.Bill) error {
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
