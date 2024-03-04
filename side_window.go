package rratelimit

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var luaSlideWindow string

type SlideWindowLimiter struct {
	client   redis.Cmdable
	service  string
	interval time.Duration
	rate     int
}

var _ Limiter = &SlideWindowLimiter{}
var _ Creator = NewSlideWindowLimiter

func NewSlideWindowLimiter(client redis.Cmdable, service string,
	interval time.Duration, rate int) Limiter {
	return &SlideWindowLimiter{
		client:   client,
		service:  service,
		interval: interval,
		rate:     rate,
	}
}
func (f *SlideWindowLimiter) Allow(ctx context.Context) (bool, error) {
	return f.client.Eval(ctx, luaFixWindow, []string{f.service}, f.interval, f.rate).Bool()
}
