package endpoint

import (
	"context"
	"folder/internal/entity"
	"folder/internal/service"

	"github.com/go-kit/kit/endpoint"
)

func MakeClientEndpoints(s service.ClientService, endpoints map[string]endpoint.Endpoint) map[string]endpoint.Endpoint {
	endpoints["FindAllClientEndpoint"] = makeFindAllClientEndpoint(s)
	endpoints["GetByIDClientEndpoint"] = makeGetByIDClientEndpoint(s)
	endpoints["CreateClientEndpoint"] = makeCreateClientEndpoint(s)
	endpoints["DeleteClientEndpoint"] = makeDeleteClientEndpoint(s)
	return endpoints
}

func makeFindAllClientEndpoint(s service.ClientService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		client, e := s.FindAll(ctx)
		return repository.FindAllClientResponse{Clients: client, Err: e}, nil
	}
}

func makeGetByIDClientEndpoint(s service.ClientService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(entity.GetByIDClientRequest)
		client, e := s.GetByID(ctx, req.ID)
		return entity.GetByIDClientResponse{Client: client, Err: e}, nil
	}
}

func makeCreateClientEndpoint(s service.ClientService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(entity.CreateClientRequest)
		e := s.Create(ctx, req.Client)
		return entity.CreateClientResponse{Err: e}, nil
	}
}

func makeDeleteClientEndpoint(s service.ClientService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(entity.DeleteClientRequest)
		e := s.Delete(ctx, req.ID)
		return entity.DeleteClientResponse{Err: e}, nil
	}
}
