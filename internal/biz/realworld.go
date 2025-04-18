package biz

import (
	"context"

	v1 "kratos-realworld/api/realworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type Realworld struct {
	Hello string
}

// GreeterRepo is a Greater repo.
type RealworldRepo interface {
	Save(context.Context, *Realworld) (*Realworld, error)
	Update(context.Context, *Realworld) (*Realworld, error)
	FindByID(context.Context, int64) (*Realworld, error)
	ListByHello(context.Context, string) ([]*Realworld, error)
	ListAll(context.Context) ([]*Realworld, error)
}

// GreeterUsecase is a Greeter usecase.
type RealworldUsecase struct {
	repo RealworldRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewRealworldUsecase(repo RealworldRepo, logger log.Logger) *RealworldUsecase {
	return &RealworldUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *RealworldUsecase) CreateRealworld(ctx context.Context, g *Realworld) (*Realworld, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
