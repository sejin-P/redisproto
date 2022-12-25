package redis

import (
	"github.com/go-redis/redis/v8"
)

func New(address, password string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
	})
}
