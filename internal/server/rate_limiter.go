package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metsawyr/abios-api/internal/config"
	"github.com/redis/go-redis/v9"
)

func newRateLimiter(config *config.Config, redisClient *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := fmt.Sprintf("%v:%v", ctx.ClientIP(), ctx.Request.URL.Path)
		err := redisClient.SetNX(ctx, key, config.App.RateLimit.Requests, time.Second*time.Duration(config.App.RateLimit.Seconds)).Err()
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		value, err := redisClient.Get(ctx, key).Result()
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		remaining, _ := strconv.ParseInt(value, 10, 64)
		if remaining == 0 {
			ctx.AbortWithError(http.StatusTooManyRequests, fmt.Errorf("quota exceeded"))
			return
		}

		err = redisClient.DecrBy(ctx, key, 1).Err()
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
}
