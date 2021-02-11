package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"folder/internal/entity"

	"github.com/gorilla/mux"
)

func DecodeFindAllDatosRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return entity.FindAllProductRequest{}, nil
}

func DecodeCreateDatosRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req entity.CreateDatosRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Product); e != nil {
		return nil, e
	}
	return req, nil
}
func DecodeDeleteDatosRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrMissingRequiredArguments
	}
	return entity.DeleteDatosRequest{ID: id}, nil
}

func EncodeDatosResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		EncodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
