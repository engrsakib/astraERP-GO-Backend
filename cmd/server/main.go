package main

import (
    "log"

    "github.com/engrsakib/erp-system/internal/config"
    "github.com/engrsakib/erp-system/internal/db"
    "github.com/engrsakib/erp-system/internal/http/routes"
    "github.com/engrsakib/erp-system/internal/models"
)

// @title           ERP System API
// @version         1.0
// @description     ERP backend with Go, Gin, GORM, PostgreSQL, Redis.
// @host      localhost:8080
// @BasePath  /
func main() {
    cfg := config.LoadConfig()

    pg := db.NewPostgres(cfg.DatabaseURL)

    if err := pg.AutoMigrate(&models.User{}); err != nil {
        log.Fatalf("auto migrate failed: %v", err)
    }

    redisClient := db.NewRedis(cfg.RedisAddr, cfg.RedisPass, cfg.RedisDB)
    _ = redisClient // future use (cache, session, etc.)

    r := routes.NewRouter(pg, redisClient)

    log.Printf("Server running on %s", cfg.AppPort)
    if err := r.Run(cfg.AppPort); err != nil {
        log.Fatalf("server failed: %v", err)
    }
}
