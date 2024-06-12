package repository

import (
	"github.com/jmoiron/sqlx"
)

type Link interface {
	AddLink(originalURL, shortURL string) (int, error)
	GetShortByOriginalURL(originalURL string) (string, error)
}

type Repository struct {
	Link
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Link: NewLinkPostgres(db),
	}
}
