package test

import (
	"backend/domain/model"
	"backend/repository"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateLink(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	repo := repository.NewLinkRepository(mockDB)
	link := &model.Link{
		ID:        "123",
		LongURL:   "https://example.com",
		ShortURL:  "exmpl",
		CreatedAt: time.Now(),
	}

	mock.ExpectExec(`INSERT INTO links \(id, long_url, short_url, created_at\) VALUES \(\$1, \$2, \$3, \$4\)`).
		WithArgs(link.ID, link.LongURL, link.ShortURL, link.CreatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Create(link)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetByShortURL(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	repo := repository.NewLinkRepository(mockDB)
	shortURL := "exmpl"
	expectedLink := &model.Link{
		ID:        "123",
		LongURL:   "https://example.com",
		ShortURL:  shortURL,
		CreatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "long_url", "short_url", "created_at"}).
		AddRow(expectedLink.ID, expectedLink.LongURL, expectedLink.ShortURL, expectedLink.CreatedAt)
	mock.ExpectQuery(`SELECT id, long_url, short_url, created_at FROM links WHERE short_url = \$1`).
		WithArgs(shortURL).
		WillReturnRows(rows)

	link, err := repo.GetByShortURL(shortURL)
	assert.NoError(t, err)
	assert.Equal(t, expectedLink, link)
	assert.NoError(t, mock.ExpectationsWereMet())
}
