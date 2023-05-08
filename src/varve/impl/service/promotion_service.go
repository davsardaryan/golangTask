package service

import (
	"context"
	"github.com/go-redis/redis/v8"
	"golangTask/src/varve/impl/entities"
	"time"
)

// Store data in Redis DB.
func save(promotion entities.Promotion, rdb *redis.Client) {

	rdb.Set(context.Background(), promotion.ID, promotion, 0)

}

// GetById read data from Redis DB.
func GetById(id string, rdb *redis.Client) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	val, err := rdb.Get(ctx, id).Result()

	return val, err
}
