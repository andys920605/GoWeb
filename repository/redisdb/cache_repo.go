package redisdb

import (
	models_svc "GoWeb/models/service"
	rep_interface "GoWeb/repository/interface"
	"context"
	"encoding/json"
	"time"

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
func NewCacheRepository(redisClient *redis.Client) rep_interface.ICacheRep {
	return &CacheRepository{
		client: redisClient,
	}
}

// region login token cache
// get login token cache
func (cache *CacheRepository) GetTokenByIDCtx(ctx context.Context, key string) (*models_svc.Claims, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CacheRepository.GetTokenByIDCtx")
	defer span.Finish()
	dataBytes, err := cache.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "CacheRepository.GetTokenByIDCtx.redisClient.Get")
	}
	data := &models_svc.Claims{}
	if err = json.Unmarshal(dataBytes, data); err != nil {
		return nil, errors.Wrap(err, "CacheRepository.GetTokenByIDCtx.json.Unmarshal")
	}
	return data, nil
}

// set login token cache
func (cache *CacheRepository) SetTokenCtx(ctx context.Context, key string, seconds int, item *models_svc.Claims) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CacheRepository.SetTokenCtx")
	defer span.Finish()
	dataBytes, err := json.Marshal(item)
	if err != nil {
		return errors.Wrap(err, "CacheRepository.SetTokenCtx.json.Marshal")
	}
	if err = cache.client.Set(ctx, key, dataBytes, time.Second*time.Duration(seconds)).Err(); err != nil {
		return errors.Wrap(err, "CacheRepository.SetTokenCtx.redis.Set")
	}
	return nil
}

// endregion

// region email random code cache
// get email random code cache
func (cache *CacheRepository) GetEmailCodeCtx(ctx context.Context, key string) ([]string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CacheRepository.GetEmailCodeCtx")
	defer span.Finish()
	stringModel, err := cache.client.SMembers(ctx, key).Result()
	if err != nil {
		return nil, errors.Wrap(err, "CacheRepository.GetEmailCodeCtx.redisClient.Get")
	}
	if len(stringModel) == 0 {
		return nil, nil
	}
	return stringModel, nil
}

// set email random code cache
func (cache *CacheRepository) SetEmailCodeCtx(ctx context.Context, key string, seconds int, item string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CacheRepository.SetEmailCodeCtx")
	defer span.Finish()
	if err := cache.client.SAdd(ctx, key, item).Err(); err != nil {
		return errors.Wrap(err, "CacheRepository.SetEmailCodeCtx.redis.Set")
	}
	cache.client.Do(ctx, "EXPIRE", key, seconds)
	//cache.client.Expire(ctx, key, time.Second*time.Duration(seconds)).Result()
	return nil
}

// endregion

// delete cache
func (cache *CacheRepository) DeleteCtx(ctx context.Context, key string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CacheRepository.DeleteCtx")
	defer span.Finish()
	if err := cache.client.Del(ctx, key).Err(); err != nil {
		return errors.Wrap(err, "CacheRepository.DeleteCtx.redis.Del")
	}
	return nil
}
