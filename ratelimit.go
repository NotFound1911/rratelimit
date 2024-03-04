package rratelimit

import (
	"github.com/redis/go-redis/v9"
	"time"
)

type RateLimit struct {
	Limiter    Limiter
	valCreator Creator
}
type RateLimitOption func(*RateLimit) error

func NewRateLimit(client redis.Cmdable, service string,
	interval time.Duration, rate int, opts ...RateLimitOption) *RateLimit {
	r := &RateLimit{valCreator: NewFixWindowLimiter}
	for _, opt := range opts {
		if err := opt(r); err != nil {
			panic(err)
		}
	}
	r.Limiter = r.valCreator(client, service, interval, rate)
	return r
}
func WithSideWindow() RateLimitOption {
	return func(limit *RateLimit) error {
		limit.valCreator = NewSlideWindowLimiter
		return nil
	}
}
