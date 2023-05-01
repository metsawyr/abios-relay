package resources

import (
	"github.com/metsawyr/abios-api/internal/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(config *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Redis.Uri,
		Password: config.Redis.Password,
		DB:       config.Redis.Database,
	})
}
