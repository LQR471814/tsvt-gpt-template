package main

import (
	"context"
	"tsvt-gp-template/server/api"

	"github.com/golang/protobuf/ptypes/empty"
)

type APIServer struct {
	api.UnimplementedServiceServer
}

func (s *APIServer) Echo(_ context.Context, _ *empty.Empty) (*api.Message, error) {
	return &api.Message{
		Content: "Hello world!",
	}, nil
}
