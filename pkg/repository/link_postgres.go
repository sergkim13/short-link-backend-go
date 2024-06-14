package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)


type LinkPostgres struct {
	db *sqlx.DB
}

func (r *LinkPostgres) AddLink(originalURL, shortURL string) (int, error) {
	var linkID int

	query := fmt.Sprintf("INSERT INTO %s (original, short) values ($1, $2) RETURNING id", linksTable)
	row := r.db.QueryRow(query, originalURL, shortURL)

	if err := row.Scan(&linkID); err != nil {
		return 0, fmt.Errorf("failed to INSERT link with original url - %s, and short url %s: %w", originalURL, shortURL, err)
	}

	return linkID, nil
}

func (r *LinkPostgres) GetShortByOriginalURL(originalURL string) (string, error) {
	var shortURL string

	query := fmt.Sprintf("SELECT short FROM %s WHERE original = $1", linksTable)
	err := r.db.Get(&shortURL, query, originalURL)

	if err != nil {
		return "", fmt.Errorf("failed to SELECT link with original url - %s: %w", originalURL, err)
	}

	return shortURL, nil
}

func (r *LinkPostgres) GetOriginalByShortURL(shortURL string) (string, error) {
	var originalURL string

	query := fmt.Sprintf("SELECT original FROM %s WHERE short = $1", linksTable)
	err := r.db.Get(&originalURL, query, shortURL)

	if err != nil {
		return "", fmt.Errorf("failed to SELECT link with short url - %s: %w", shortURL, err)
	}

	return originalURL, nil
}


func NewLinkPostgres(db *sqlx.DB) *LinkPostgres {
	return &LinkPostgres{db: db}
}
