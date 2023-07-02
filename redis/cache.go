package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/vishal1132/bootstrap-go-web/config"
	"github.com/vishal1132/bootstrap-go-web/utils"
	"go.uber.org/zap"
)

func InitRedis(ctx context.Context, redisConfig *config.RedisConfig) *redis.Client {
	zap.L().Info("Connecting to redis")
	address := redisConfig.Host + ":" + redisConfig.Port
	cfg := redis.Options{
		Addr: address,
		DB:   redisConfig.DbId,
	}

	redisClient := redis.NewClient(&cfg)
	_ = utils.Must(redisClient.Ping(ctx).Result())

	zap.L().Info("Connected to redis")
	return redisClient
}
