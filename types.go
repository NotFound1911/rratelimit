package rratelimit

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Limiter interface {
	Allow(ctx context.Context) (bool, error) // 是否限流
}
type Creator func(client redis.Cmdable, service string,
	interval time.Duration, rate int) Limiter
