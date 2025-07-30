package middleware

import (
	"complaint-service/config"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func RateLimitRedisMiddleware(limit int, duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := fmt.Sprintf("ratelimit:%s", ip)
		log.Printf("RateLimitRedisMiddleware: Checking rate limit for IP: " + ip)
		val, err := config.RedisClient.Get(config.Ctx, key).Int()
		if err != nil && err.Error() != "redis: nil" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Redis error"})
			return
		}

		if val >= limit {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests. Please wait.",
			})
			return
		}

		pipe := config.RedisClient.TxPipeline()
		pipe.Incr(config.Ctx, key)
		pipe.Expire(config.Ctx, key, duration)
		_, _ = pipe.Exec(config.Ctx)

		c.Next()
	}
}
