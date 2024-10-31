package main

import (
	"os"

	"urlshortener/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	log := logger.New("info")
	router := gin.Default()

	port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default to port 8080 if not set
    }
    log.Infof("Starting server on port %s", "8080")
    if err := router.Run(":" + port); err != nil {
        log.Fatalf("failed to start server: %v", err)
    }
}
