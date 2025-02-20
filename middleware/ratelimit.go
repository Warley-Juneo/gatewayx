package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/go-redis/redis/v8"
    "context"
    "time"
)

func RateLimiter(redisClient *redis.Client, limit int, window time.Duration) gin.HandlerFunc {
    return func(c *gin.Context) {
        ip := c.ClientIP()
        key := "rate_limit:" + ip

        count, err := redisClient.Incr(context.Background(), key).Result()
        if err != nil {
            c.AbortWithStatusJSON(500, gin.H{"error": "Erro interno"})
            return
        }

        if count == 1 {
            redisClient.Expire(context.Background(), key, window)
        }

        if count > int64(limit) {
            c.AbortWithStatusJSON(429, gin.H{"error": "Limite de requisições excedido"})
            return
        }

        c.Next()
    }
}