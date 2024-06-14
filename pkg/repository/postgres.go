package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sergkim13/short-link-backend-go/configs"
)


const linksTable = "links"

func NewPostgresDB(cfg configs.Config) (*sqlx.DB, error) {
	database, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = database.Ping()

	if err != nil {
		return nil, err
	}

	return database, nil
}
