package storage

import (
	"context"
	"github.com/Imomali1/gophermart/internal/apps/gophermart/config"
)

type IStorage interface {
	Help()
}

func New(ctx context.Context, cfg config.Config) (IStorage, error) {
	return newDBStorage(ctx, cfg.DatabaseDSN)
}
