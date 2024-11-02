package handler

import (
	"net/http"
	"urlshortener/internal/app/services"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	URLService *services.URLService
}

func NewURLHandler(urlService *services.URLService) *URLHandler {
    return &URLHandler{URLService: urlService}
}

func (h *URLHandler) ShortenURLHandler(c *gin.Context) {
    var req struct {
        OriginalURL string `json:"original_url" binding:"required,url"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    url, err := h.URLService.ShortenURL(req.OriginalURL)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to shorten URL"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "short_url": url.ShortCode,
        "original_url": url.OriginalURL,
    })
}

func (h *URLHandler) RedirectHandler(c *gin.Context) {
    shortCode := c.Param("shortCode")

    url, err := h.URLService.GetOriginalURL(shortCode)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
        return
    }

    go h.URLService.TrackClick(shortCode)

    c.Redirect(http.StatusFound, url)
}
