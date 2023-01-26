package main

import (
	"context"
	api "example/kitex_gen/api"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the EchoImpl interface.
func (s *EchoImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	return &api.Response{
		Sum:     req.A + req.B,
		Message: req.Message,
	}, nil
}
