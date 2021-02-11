package endpoint

import (
	"context"

	"folder/internal/entity"
	"folder/internal/service"

	"github.com/go-kit/kit/endpoint"
)

func MakeDatosEndpoints(s service.DatosService, endpoints map[string]endpoint.Endpoint) map[string]endpoint.Endpoint {
	endpoints["FindAllDatosEndpoint"] = makeFindAllEndpoint(s)
	endpoints["CreateDatosEndpoint"] = makeCreateEndpoint(s)
	endpoints["DeleteDatosEndpoint"] = makeDeleteDatosEndpoint(s)
	return endpoints
}
func makeFindAllEndpoint(s service.DatosService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		datos, e := s.FindAll(ctx)
		return entity.FindAllDatosResponse{TDatos: datos, Err: e}, nil
	}
}

func makeCreateEndpoint(s service.DatosService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(entity.CreateRequest)
		e := s.Create(ctx, req.Productos)
		return entity.CreateDatosResponse{Err: e}, nil
	}
}
func makeDeleteDatosEndpoint(s service.DatosService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(entity.DeleteOrderRequest)
		e := s.Delete(ctx, req.ID)
		return entity.DeleteDatosResponse{Err: e}, nil
	}
}
