package storage

import "context"

type IStorage interface {
	Help()
}

func New(ctx context.Context, cfg config.Config) (IStorage, error) {
	return newDBStorage(ctx, cfg)
}
