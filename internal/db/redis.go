package db

import (
    "context"
    "log"

    "github.com/redis/go-redis/v9"
)

func NewRedis(addr, password string, db int) *redis.Client {
    rdb := redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: password,
        DB:       db,
    })

    if err := rdb.Ping(context.Background()).Err(); err != nil {
        log.Fatalf("failed to connect redis: %v", err)
    }

    return rdb
}
