/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/9/14 19:31
 */
package zcache

import (
	"context"
	"sync"
	"time"

	"github.com/go-kirito/pkg/zcache/redis"

	"github.com/go-kirito/pkg/zconfig"
	"github.com/go-kirito/pkg/zlog"
)

type ICache interface {
	Set(ctx context.Context, key string, value interface{}, expire time.Duration) error
	SetNX(ctx context.Context, key string, value interface{}, expire time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	SAdd(ctx context.Context, key string, value []interface{}, expire time.Duration) error
	SMembers(ctx context.Context, key string) ([]string, error)
	Exists(ctx context.Context, key string) bool
	ZAdd(ctx context.Context, key string, score float64, value interface{}) error
	ZIncr(ctx context.Context, key string, score float64, member string, expire time.Duration) error
	ZRank(ctx context.Context, key string, member string) int64
	ZScore(ctx context.Context, key string, member string) float64
	ZRevRange(ctx context.Context, key string, start int64, end int64) []map[string]float64
}

type Options struct {
	Driver string
	Type   string
}

var cache ICache
var once = &sync.Once{}

func NewCache() ICache {
	once.Do(func() {
		var opts Options
		if err := zconfig.UnmarshalKey("cache", &opts); err != nil {
			zlog.Fatal("read cache config err", err)
		}

		if opts.Driver == "redis" {
			cache = redis.NewCache()
		}
	})
	return cache
}
