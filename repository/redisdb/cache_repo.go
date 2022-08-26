package redisdb

import (
	rep_interface "GoWeb/repository/interface"
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type CacheRepository struct {
	client *redis.Client
}

// new cache repository instance
// @param client
// @result cache repository
func NewCacheRepository(redisClient *redis.Client) rep_interface.ICacheRepo {
	return &CacheRepository{
		client: redisClient,
	}
}

// Delete from cache
// @param context
// @param key
// @result error
func (cache *CacheRepository) DeleteCtx(ctx context.Context, key string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CacheRepository.DeleteCtx")
	defer span.Finish()
	if err := cache.client.Del(ctx, key).Err(); err != nil {
		return errors.Wrap(err, "CacheRepository.DeleteCtx.redis.Del")
	}
	return nil
}
