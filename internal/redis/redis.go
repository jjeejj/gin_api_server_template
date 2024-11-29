package redis

import (
	"context"
	"time"

	goRedis "github.com/redis/go-redis/v9"
)

// RedisConfig 配置
type RedisConfig struct {
	Addr     string
	Password string
	Db       int
	Prefix   string
}

type RedisClient struct {
	redisClient *goRedis.Client
	Prefix      string // 前缀
}

func New(config *RedisConfig) *RedisClient {
	rdb := goRedis.NewClient(&goRedis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.Db,
	})
	return &RedisClient{
		redisClient: rdb,
		Prefix:      config.Prefix,
	}
}

func (r *RedisClient) Close() {
	r.redisClient.Close()
}

// Set 设置 string
func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.redisClient.Set(ctx, r.Prefix+key, value, expiration).Err()
}

// Get 获取 string
func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return r.redisClient.Get(ctx, r.Prefix+key).Result()
}

// 删除 key
func (r *RedisClient) Del(ctx context.Context, key string) error {
	return r.redisClient.Del(ctx, r.Prefix+key).Err()
}

// 设置过期时间
func (r *RedisClient) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return r.redisClient.Expire(ctx, r.Prefix+key, expiration).Err()
}

func (r *RedisClient) ZAdd(ctx context.Context, key string, members ...goRedis.Z) error {
	return r.redisClient.ZAdd(ctx, r.Prefix+key, members...).Err()
}

func (r *RedisClient) ZRangeByScore(ctx context.Context, key string, opt *goRedis.ZRangeBy) error {
	return r.redisClient.ZRangeByScore(ctx, r.Prefix+key, opt).Err()
}

func (r *RedisClient) ZRevRangeByScore(ctx context.Context, key string, opt *goRedis.ZRangeBy) ([]string, error) {
	var result = r.redisClient.ZRevRangeByScore(ctx, r.Prefix+key, opt)
	return result.Val(), result.Err()
}
