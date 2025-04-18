package service

import (
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
