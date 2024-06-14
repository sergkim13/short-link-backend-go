package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)


type LinkPostgres struct {
	db *sqlx.DB
}

func (r *LinkPostgres) AddLink(originalURL, shortURL string) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (original, short) values ($1, $2) RETURNING id", linksTable)
	row := r.db.QueryRow(query, originalURL, shortURL)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *LinkPostgres) GetShortByOriginalURL(originalURL string) (string, error) {
	var shortURL string
	query := fmt.Sprintf("SELECT short FROM %s WHERE original = $1", linksTable)
	err := r.db.Get(&shortURL, query, originalURL)

	return shortURL, err
}

func (r *LinkPostgres) GetOriginalByShortURL(shortURL string) (string, error) {
	var originalURL string
	query := fmt.Sprintf("SELECT original FROM %s WHERE short = $1", linksTable)
	err := r.db.Get(&originalURL, query, shortURL)

	return originalURL, err
}


func NewLinkPostgres(db *sqlx.DB) *LinkPostgres {
	return &LinkPostgres{db: db}
}
