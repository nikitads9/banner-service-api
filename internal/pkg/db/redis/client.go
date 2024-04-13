package rediska

import "github.com/go-redis/redis"

func GetClient(address string, password string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})
}
