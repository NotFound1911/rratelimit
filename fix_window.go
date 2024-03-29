package rratelimit

import (
	"context"
	_ "embed"
	"github.com/redis/go-redis/v9"
	"time"
)

//go:embed lua/fix_window.lua
var luaFixWindow string

type FixWindowLimiter struct {
	client   redis.Cmdable
	service  string
	interval time.Duration
	rate     int // 阈值
}

func (f *FixWindowLimiter) Allow(ctx context.Context) (bool, error) {
	return f.client.Eval(ctx, luaFixWindow, []string{f.service}, f.interval, f.rate).Bool()
}

var _ Limiter = &FixWindowLimiter{}
var _ Creator = NewFixWindowLimiter

func NewFixWindowLimiter(client redis.Cmdable, service string,
	interval time.Duration, rate int) Limiter {
	return &FixWindowLimiter{
		client:   client,
		service:  service,
		interval: interval,
		rate:     rate,
	}
}
