package handler

import (
	"context"
	"errors"

	"github.com/HyanSource/grpctest/api"
)

type APIServer struct {
	api.UnimplementedAPIServiceServer
}

func (a *APIServer) GetUser(ctx context.Context, req *api.GetUserRequest) (*api.GetUserResponse, error) {

	if req.Id == 0 {
		return &api.GetUserResponse{}, errors.New("No id")
	}

	return &api.GetUserResponse{
		Id:    req.Id,
		Name:  "Hyan",
		Money: 777,
	}, nil
}
