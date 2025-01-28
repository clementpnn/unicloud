package repository

import (
	"backend/domain/model"
	"database/sql"
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
