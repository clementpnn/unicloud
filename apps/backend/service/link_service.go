package service

import (
	"backend/domain/model"
	"backend/repository"
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/google/uuid"
)

type LinkService interface {
	CreateShortURL(longURL string) (*model.Link, error)
	GetByShortURL(shortURL string) (*model.Link, error)
}

type LinkServiceImpl struct {
	repo *repository.LinkRepository
}

func NewLinkService(repo *repository.LinkRepository) *LinkServiceImpl {
	return &LinkServiceImpl{repo: repo}
}

func (s *LinkServiceImpl) CreateShortURL(longURL string) (*model.Link, error) {
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

func (s *LinkServiceImpl) GetByShortURL(shortURL string) (*model.Link, error) {
	return s.repo.GetByShortURL(shortURL)
}
