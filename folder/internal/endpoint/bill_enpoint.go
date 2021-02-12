package endpoint

import (
	"context"
	"folder/internal/entity"
	"folder/internal/service"

	"github.com/go-kit/kit/endpoint"
)

func MakeBillEndpoints(s service.BillService, endpoints map[string]endpoint.Endpoint) map[string]endpoint.Endpoint {
	endpoints["FindAllBillEndpoint"] = makeFindAllBillEndpoint(s)
	endpoints["GetByIDBillEndpoint"] = makeGetByIDBillEndpoint(s)
	endpoints["CreateBillEndpoint"] = makeCreateBillEndpoint(s)
	endpoints["DeleteBillEndpoint"] = makeDeleteBillEndpoint(s)
	return endpoints
}

func makeFindAllBillEndpoint(s service.BillService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		bill, e := s.FindAll(ctx)
		return repository.FindAllBillResponse{Bills: bill, Err: e}, nil
	}
}

func makeGetByIDBillEndpoint(s service.BillService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(entity.GetByIDBillRequest)
		bill, e := s.GetByID(ctx, req.ID)
		return entity.GetByIDBillResponse{Bill: bill, Err: e}, nil
	}
}

func makeCreateBillEndpoint(s service.BillService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(entity.CreateBillRequest)
		e := s.Create(ctx, req.Bill)
		return entity.CreateBillResponse{Err: e}, nil
	}
}

func makeDeleteBillEndpoint(s service.BillService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(entity.DeleteBillRequest)
		e := s.Delete(ctx, req.ID)
		return entity.DeleteBillResponse{Err: e}, nil
	}
}
