package service

import (
	"backend/domain/model"
	"backend/repository"
	"crypto/rand"
	"encoding/base64"
	"time"
	"github.com/google/uuid"
)

type LinkService struct {
	repo *repository.LinkRepository
}

func NewLinkService(repo *repository.LinkRepository) *LinkService {
	return &LinkService{repo: repo}
}

func (s *LinkService) CreateShortURL(longURL string) (*model.Link, error) {
	shortURL, err := generateShortURL()
	if err != nil {
		return nil, err
	}

	link := &model.Link{
		ID:        uuid.New().String(),
		LongURL:   longURL,
		ShortURL:  shortURL,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(link); err != nil {
		return nil, err
	}

	return link, nil
}

func generateShortURL() (string, error) {
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b)[:8], nil
}

func (s *LinkService) GetByShortURL(shortURL string) (*model.Link, error) {
	return s.repo.GetByShortURL(shortURL)
}

func (s *LinkService) GetOriginalURL(shortURL string) (string, error) {

	originalURL, err := s.repo.GetOriginalURL(shortURL)
	if err != nil {
		return "", err
	}

	return originalURL, nil
}
