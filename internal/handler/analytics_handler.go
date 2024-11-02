package handler

import (
	"net/http"
	"urlshortener/internal/app/services"

	"github.com/gin-gonic/gin"
)

type AnalyticsHandler struct {
	URLService *services.URLService
}

func NewAnalyticsHandler(urlService *services.URLService) *AnalyticsHandler {
	return &AnalyticsHandler{URLService: urlService}
}

func (h *AnalyticsHandler) GetClickStatsHandler(c *gin.Context) {
	shortCode := c.Param("shortCode")

	url, err := h.URLService.GetClickURL(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"short_url":    url.ShortCode,
		"original_url": url.OriginalURL,
		"clicks":       url.Clicks,
	})
}
