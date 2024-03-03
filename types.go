package rratelimit

import "context"

type Limiter interface {
	Allow(ctx context.Context) (bool, error) // 是否限流
}
