package main

import (
	"log"
	"time"

	"urlshortener/config"
	"urlshortener/internal/app/services"
	"urlshortener/internal/handler"
	"urlshortener/internal/repository"
	"urlshortener/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("failed to load config: %v", err)
    }

    appLog := logger.New(cfg.LogLevel)

    // permanent repo: urlRepo
    db, err := repository.NewPostgresDB(cfg.DatabaseURL)
    if err != nil {
        appLog.Fatalf("failed to connect to PostgreSQL: %v", err)
    }
    defer func() {
        sqlDB, _ := db.DB()
        sqlDB.Close()
    }()

    urlRepo := repository.NewURLRepository(db)

    // cache repo: cacheRepo
    cache := repository.NewRedisClient(cfg.RedisAddress, cfg.RedisPassword, cfg.RedisDB)
    defer cache.Close()

    cacheRepo := repository.NewCacheRepository(cache)

    urlService := services.NewURLService(urlRepo, cacheRepo, 24*time.Hour)

    // handlers
    urlHandler := handler.NewURLHandler(urlService)
    analyticsHandler := handler.NewAnalyticsHandler(urlService)

    // server
    router := gin.Default()
    port := cfg.Port
    
    router.POST("/shorten", urlHandler.ShortenURLHandler)
    router.GET("/:shortCode", urlHandler.RedirectHandler)
    router.GET("/analytics/:shortCode", analyticsHandler.GetClickStatsHandler)

    appLog.Infof("Starting server on port %s", port)
    if err := router.Run(":" + port); err != nil {
        appLog.Fatalf("failed to start server: %v", err)
    }
}
