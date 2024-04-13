package rediska

import "github.com/redis/go-redis/v9"

func GetClient(address string, password string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})
}
