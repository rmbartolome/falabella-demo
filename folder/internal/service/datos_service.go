package service

import (
	"context"

	"folder/internal/entity"

	"github.com/go-kit/kit/log"
)

type DatosService interface {
	FindAll(ctx context.Context) ([]entity.Product, error)
	Create(ctx context.Context, dat entity.Product) error
	Delete(ctx context.Context, id string) error
}

type datosService struct {
	repository entity.DatosRepository
	logger     log.Logger
	tlogger    LogsService
}

func NewDatosService(repository entity.DatosRepository, logger log.Logger, tlogger LogsService) DatosService {
	return &datosService{
		repository: repository,
		logger:     logger,
		tlogger:    tlogger,
	}
}

func (s *datosService) FindAll(ctx context.Context) ([]entity.Product, error) {
	produ, err := s.repository.FindAll(ctx)
	if err != nil {
		s.logger.Log("error:", err)
		return nil, err
	}
	s.logger.Log("findall:", "success")
	return produ, nil
}

func (s *datosService) Create(ctx context.Context, produ entity.Product) error {
	produ.Prepare()
	err := produ.Validate()
	if err != nil {
		s.logger.Log("error:", err)
		return err
	}

	if err := s.repository.Create(ctx, produ); err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("create:", "success")

	go s.tlogger.SaveLog(TLog{
		ServiceName: "Product",
		Caller:      "Products->Create",
		Event:       "POST",
		Extra:       "Create new products.",
	})
	return nil
}
func (s *datosService) Delete(ctx context.Context, id string) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("delete:", "success")
	return nil
}
