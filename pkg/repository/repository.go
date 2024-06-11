package repository

import "github.com/jmoiron/sqlx"

type Link interface {

}

type Repository struct {
	Link
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
