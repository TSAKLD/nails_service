package repository

import (
	"database/sql"
	"github.com/go-redis/redis/v9"
)

type Repository struct {
	db *sql.DB
	c  *redis.Client
}
