package repository

import (
	"backend/domain/model"
	"database/sql"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	linkCreationCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "link_creation_total",
		Help: "Total number of created links",
	})
)

type LinkRepository struct {
	db *sql.DB
}

func NewLinkRepository(db *sql.DB) *LinkRepository {
	return &LinkRepository{db: db}
}

func (r *LinkRepository) Create(link *model.Link) error {
	query := `
        INSERT INTO links (id, long_url, short_url, created_at)
        VALUES ($1, $2, $3, $4)
    `
	_, err := r.db.Exec(query, link.ID, link.LongURL, link.ShortURL, link.CreatedAt)
	return err
}

func (r *LinkRepository) GetByShortURL(shortURL string) (*model.Link, error) {
	link := &model.Link{}
	query := `
        SELECT id, long_url, short_url, created_at
        FROM links WHERE short_url = $1
    `
	err := r.db.QueryRow(query, shortURL).Scan(
		&link.ID, &link.LongURL, &link.ShortURL, &link.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return link, nil
}

func (r *LinkRepository) CreateShortURL(url string) (*model.Link, error) {
	link := &model.Link{}
	linkCreationCounter.Inc()

	return link, nil
}

func (r *LinkRepository) GetOriginalURL(shortURL string) (string, error) {
	var originalURL string

	query := `
        SELECT long_url 
        FROM links 
        WHERE short_url = $1
        LIMIT 1
    `

	err := r.db.QueryRow(query, shortURL).Scan(&originalURL)

	if err == sql.ErrNoRows {
		return "", fmt.Errorf("URL not found")
	}

	if err != nil {
		return "", fmt.Errorf("database error: %v", err)
	}

	return originalURL, nil
}
