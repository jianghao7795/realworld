package service

import (
	"context"

	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/biz"
)

// GreeterService is a greeter service.
type RealworldService struct {
	v1.UnimplementedRealworldServer

	uc *biz.RealworldUsecase
}

// NewGreeterService new a greeter service.
func NewRealworldService(uc *biz.RealworldUsecase) *RealworldService {
	return &RealworldService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *RealworldService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateRealworld(ctx, &biz.Realworld{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
