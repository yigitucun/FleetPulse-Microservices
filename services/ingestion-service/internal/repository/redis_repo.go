package repository

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(addr string) *RedisRepository {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &RedisRepository{client: rdb}
}

func (r *RedisRepository) SaveLastPosition(truckID string, data string) error {
	return r.client.Set(ctx, "truck:"+truckID, data, 10*time.Minute).Err()
}

func (r *RedisRepository) GetLastPosition(truckID string) (string, error) {
	return r.client.Get(ctx, "truck:"+truckID).Result()
}