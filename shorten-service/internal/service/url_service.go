package service

import (
	"shorten-service/internal/repository"
)

type URLService struct {
	repo *repository.URLRepository
}

func NewURLService(repo *repository.URLRepository) *URLService {
	return &URLService{repo}
}

func (s *URLService) ShortenURL(longURL string) (string, error) {

	// First deterministic key

	existingShortKey, err := s.repo.GetShortKeyByLongURL(longURL)
	if err == nil && existingShortKey != "" {
		return existingShortKey, nil
	}

	shortKey := createShortKey(longURL)

	// Check if key exists
	existingURL, _ := s.repo.GetOriginalURL(shortKey)
	if existingURL != "" && existingURL != longURL {
		// Collision -> fallback to salted key
		shortKey = createSaltedShortKey(longURL)
	}

	// Save
	err = s.repo.SaveShortURL(shortKey, longURL)
	if err != nil {
		return "", err
	}

	return shortKey, nil
}

func (s *URLService) GetOriginalURL(shortKey string) (string, error) {
	return s.repo.GetOriginalURL(shortKey)
}
