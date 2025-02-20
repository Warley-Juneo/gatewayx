package storage

import (
    "github.com/go-redis/redis/v8"
    "context"
)

func NewRedisClient(addr string) *redis.Client {
    return redis.NewClient(&redis.Options{
        Addr: addr,
    })
}

func PingRedis(client *redis.Client) error {
    _, err := client.Ping(context.Background()).Result()
    return err
}