package database

import (
	"github.com/Yakov-Varnaev/ft/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New(cfg config.PostgresConfig) *sqlx.DB {
	return sqlx.MustConnect("postgres", cfg.AsConnString())
}
