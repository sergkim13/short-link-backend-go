package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)


type LinkPostgres struct {
	db *sqlx.DB
}

func NewLinkPostgres(db *sqlx.DB) *LinkPostgres {
	return &LinkPostgres{db: db}
}

func (r *LinkPostgres) AddLink(originalURL, shortURL string) (string, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (original, short) values ($1, $2) RETURNING id", linksTable)
	row := r.db.QueryRow(query, originalURL, shortURL)

	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return "Ok", nil
}
