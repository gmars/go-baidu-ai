package common

import (
	"context"
)

type Cache interface {
	SetData(ctx context.Context, key string, value string, expiresSeconds int64) error
	GetData(ctx context.Context, key string) (string, error)
	DelData(ctx context.Context, key string) error
}
