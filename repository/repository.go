package repository

import (
	"database/sql"
	"github.com/go-redis/redis/v9"
)

type Repository struct {
	db *sql.DB
	c  *redis.Client
}

func New(database *sql.DB, client *redis.Client) *Repository {
	return &Repository{
		db: database,
		c:  client,
	}
}
