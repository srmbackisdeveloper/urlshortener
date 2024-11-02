package services

import "urlshortener/internal/model"

func (s *URLService) TrackClick(shortCode string) error {
	return s.URLRepo.IncrementClickCount(shortCode)
}

func (s *URLService) GetClickURL(shortCode string) (*model.URL, error) {
	url, err := s.URLRepo.GetURLByShortCode(shortCode)
	if err != nil {
		return nil, err
	}
	return url, nil
}