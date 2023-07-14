/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/9/14 19:37
 */
package redis

import (
	"context"
	"errors"
	"time"

	"github.com/go-kirito/pkg/zconfig"
	"github.com/go-kirito/pkg/zlog"
	"github.com/go-redis/redis/v8"
)

type Options struct {
	Addr     string
	Password string
	Db       int
}

type cache struct {
	rdb *redis.Client
}

func NewCache() *cache {
	var opts Options

	if err := zconfig.UnmarshalKey("redis", &opts); err != nil {
		zlog.Fatal("redis error:", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     opts.Addr,
		Password: opts.Password, // no password set
		DB:       opts.Db,       // use default DB
	})

	return &cache{
		rdb: rdb,
	}
}

func (c cache) Set(ctx context.Context, key string, value interface{}, expire time.Duration) error {
	return c.rdb.Set(ctx, key, value, expire).Err()
}

func (c cache) SetNX(ctx context.Context, key string, value interface{}, expire time.Duration) (bool, error) {
	return c.rdb.SetNX(ctx, key, value, expire).Result()
}

func (c cache) Get(ctx context.Context, key string) (string, error) {
	val, err := c.rdb.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return val, nil
}

func (c cache) Del(ctx context.Context, key string) error {
	return c.rdb.Del(ctx, key).Err()
}

func (c cache) SAdd(ctx context.Context, key string, value []interface{}, expire time.Duration) error {
	_, err := c.rdb.SAdd(ctx, key, value...).Result()
	if err != nil {
		return err
	}

	ok, err := c.rdb.Expire(ctx, key, expire).Result()
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("设置过期时间失败")
	}

	return nil
}

func (c cache) SMembers(ctx context.Context, key string) ([]string, error) {
	members, err := c.rdb.SMembers(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}

	return members, nil
}

func (c cache) Exists(ctx context.Context, key string) bool {
	num, err := c.rdb.Exists(ctx, key).Result()
	if err != nil {
		zlog.Error("redis exists err:", err.Error())
		return false
	}

	return num > 0
}

func (c cache) ZAdd(ctx context.Context, key string, score float64, value interface{}) error {
	if _, err := c.rdb.ZAdd(ctx, key, &redis.Z{
		Score:  score,
		Member: value,
	}).Result(); err != nil {
		return err
	}

	return nil
}

func (c cache) ZIncr(ctx context.Context, key string, score float64, member string, expire time.Duration) error {
	count, err := c.rdb.ZCard(ctx, key).Result()
	if err != nil {
		return err
	}

	if _, err = c.rdb.ZIncrBy(ctx, key, score, member).Result(); err != nil {
		return err
	}

	if count == 0 {
		ok, err := c.rdb.Expire(ctx, key, expire).Result()
		if err != nil {
			return err
		}

		if !ok {
			return errors.New("设置过期时间失败")
		}
	}

	return nil
}

func (c cache) ZRank(ctx context.Context, key string, member string) int64 {
	return c.rdb.ZRevRank(ctx, key, member).Val()
}

func (c cache) ZScore(ctx context.Context, key string, member string) float64 {
	return c.rdb.ZScore(ctx, key, member).Val()
}

func (c cache) ZRevRange(ctx context.Context, key string, start int64, end int64) []map[string]float64 {
	list := c.rdb.ZRevRangeWithScores(ctx, key, start, end).Val()
	if list != nil {
		result := make([]map[string]float64, len(list))
		for i, v := range list {
			result[i] = map[string]float64{v.Member.(string): v.Score}
		}
		return result
	}

	return nil
}
