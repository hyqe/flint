package storage

import "context"

type KV interface {
	Get(ctx context.Context, key string)
}
