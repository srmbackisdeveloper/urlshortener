package services

import (
	"time"
	"urlshortener/internal/model"
	"urlshortener/internal/repository"

	"github.com/google/uuid"
)

type URLService struct {
	URLRepo *repository.URLRepository
	CacheRepo *repository.CacheRepository
	CacheTTL time.Duration // time to live
}

func NewURLService(urlRepo *repository.URLRepository, cacheRepo *repository.CacheRepository, cacheTTL time.Duration) *URLService {
	return &URLService{
        URLRepo:   urlRepo,
        CacheRepo: cacheRepo,
        CacheTTL:  cacheTTL,
    }
}

func (s *URLService) ShortenURL(originalURL string) (*model.URL, error) {
	shortCode := generateShortCode()

	url, err := s.URLRepo.SaveURL(originalURL, shortCode) // db
	if err != nil {
        return nil, err
    }

	err = s.CacheRepo.SetCachedURL(shortCode, originalURL, s.CacheTTL)
	if err != nil {
		return nil, err
	}

	return url, nil
}

func (s *URLService) GetOriginalURL(shortCode string) (string, error) {
	cachedURL, err := s.CacheRepo.GetCachedURL(shortCode)
	if err != nil {
        return "", err
    }

	if cachedURL != "" { // cache hit!
        return cachedURL, nil
    }
	
	url, err := s.URLRepo.GetURLByShortCode(shortCode) // cache miss
    if err != nil {
        return "", err
    }

	err = s.CacheRepo.SetCachedURL(shortCode, url.OriginalURL, s.CacheTTL) 
	if err != nil {
        return "", err
    }

	return url.OriginalURL, nil
}

func generateShortCode() string {
    return uuid.New().String()[:8]
}