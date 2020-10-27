package server

import (
	"context"
	"github/Ko4s/grpc-intro/hello"
)

//Service grpc service
type Service struct {
	hello.UnimplementedHelloServiceServer
}

//NewService return grpcSerivce
func NewService() *Service {
	return &Service{}
}

//SayHello reponse
func (s *Service) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {

	msg := req.GetMessage()

	return &hello.HelloResponse{
		Message: "Hello " + msg,
	}, nil
}
