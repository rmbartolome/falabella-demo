package handler

import (
	"context"
	"encoding/json"
	"folder/internal/service"
	"time"

	"folder/internal/entity"

	"github.com/go-kit/kit/log"
)

type LoggingDatosServiceMiddleware func(s service.DatosService) service.DatosService

type LoggingDatosServiceMiddleware struct {
	service.DatosService
	tlogger service.LogsService
	logger  log.Logger
}

func NewLoggingDatosServiceMiddleware(logger log.Logger, tlogger service.LogsService) LoggingDatosServiceMiddleware {
	return func(next service.DatosService) service.DatosService {
		return &loggingDatosServiceMiddleware{next, tlogger, logger}
	}
}

func (mw *loggingDatosServiceMiddleware) FindAll(ctx context.Context) ([]entity.Product, error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "FindAllProduct", "took", time.Since(begin))
	}(time.Now())

	collection, err := mw.DatosService.FindAll(ctx)
	tlog := service.TLog{
		ServiceName: "PRODUCT",
		Caller:      "Datos->FindAll",
		Event:       "GET",
		Extra:       "Find all products. " + time.Now().String(),
	}

	if err != nil {
		extra, _ := json.Marshal(err)
		tlog.Extra = string(extra)
	}
	go mw.tlogger.SaveLog(tlog)

	return collection, err
}

func (mw *loggingDatosServiceMiddleware) Create(ctx context.Context, product entity.Product) error {
	defer func(begin time.Time) {
		mw.logger.Log("method", "Create", "took", time.Since(begin))
	}(time.Now())

	err := mw.DatosService.Create(ctx, product)
	tlog := service.TLog{
		ServiceName: "PRODUCT",
		Caller:      "Datos->FindAll",
		Event:       "POST",
		Extra:       "Create new product. " + product.Name,
	}

	if err != nil {
		extra, _ := json.Marshal(err)
		tlog.Extra = string(extra)
	}
	go mw.tlogger.SaveLog(tlog)

	return err
}

func (mw *loggingDatosServiceMiddleware) Delete(ctx context.Context, id string) error {
	defer func(begin time.Time) {
		mw.logger.Log("method", "Delete", "took", time.Since(begin))
		mw.tlogger.SaveLog(service.TLog{
			ServiceName: "PRODUCT",
			Caller:      "Datos->FindAll",
			Event:       "DELETE",
			Extra:       "Delete product by ID.",
		})
	}(time.Now())

	err := mw.DatosService.Delete(ctx, id)
	tlog := service.TLog{
		ServiceName: "PRODUCT",
		Caller:      "Datos->FindAll",
		Event:       "DELETE",
		Extra:       "Delete product by ID. " + id,
	}

	if err != nil {
		extra, _ := json.Marshal(err)
		tlog.Extra = string(extra)
	}
	go mw.tlogger.SaveLog(tlog)

	return err
}
