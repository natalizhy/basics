package adder

import (
	"github.com/natalizhy/basics/grpcadder/pkg/api"
	"context"
)

type GRPCServer struct {
}

func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
