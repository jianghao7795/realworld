package data

import (
	"context"

	"kratos-realworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type realworldRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewRealworldRepo(data *Data, logger log.Logger) biz.RealworldRepo {
	return &realworldRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *realworldRepo) Login(ctx context.Context, g *biz.Realworld) (*biz.Realworld, error) {
	err := r.data.db.Create(g).Error
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (r *realworldRepo) Update(ctx context.Context, g *biz.Realworld) (*biz.Realworld, error) {
	return g, nil
}

func (r *realworldRepo) FindByID(context.Context, int64) (*biz.Realworld, error) {
	return nil, nil
}

func (r *realworldRepo) ListByHello(context.Context, string) ([]*biz.Realworld, error) {
	return nil, nil
}

func (r *realworldRepo) ListAll(context.Context) ([]*biz.Realworld, error) {
	return nil, nil
}
