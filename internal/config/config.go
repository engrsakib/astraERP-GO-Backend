package config

import (
    "log"
    "os"
    "strconv"

    "github.com/joho/godotenv"
)

type Config struct {
    AppPort     string
    DatabaseURL string
    RedisAddr   string
    RedisPass   string
    RedisDB     int
}

func LoadConfig() *Config {
    _ = godotenv.Load()

    return &Config{
        AppPort:     getEnv("APP_PORT", ":8080"),
        DatabaseURL: mustEnv("DATABASE_URL"),
        RedisAddr:   getEnv("REDIS_ADDR", "localhost:6379"),
        RedisPass:   getEnv("REDIS_PASSWORD", ""),
        RedisDB:     getEnvInt("REDIS_DB", 0),
    }
}

func mustEnv(key string) string {
    val := os.Getenv(key)
    if val == "" {
        log.Fatalf("missing required env: %s", key)
    }
    return val
}

func getEnv(key, def string) string {
    val := os.Getenv(key)
    if val == "" {
        return def
    }
    return val
}

func getEnvInt(key string, def int) int {
    val := os.Getenv(key)
    if val == "" {
        return def
    }
    i, err := strconv.Atoi(val)
    if err != nil {
        return def
    }
    return i
}
