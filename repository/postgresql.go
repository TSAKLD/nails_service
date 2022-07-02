package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v9"
	"log"
	"nails/bootstrap"
	"nails/entity"
)

func New(database *sql.DB, client *redis.Client) *Repository {
	return &Repository{
		db: database,
		c:  client,
	}
}

func DBConnect(c *bootstrap.Config) (*sql.DB, error) {
	info := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)

	db, err := sql.Open("postgres", info)
	if err != nil {
		return nil, err
	}

	return db, db.Ping()
}

func (r *Repository) Records() ([]entity.Record, error) {
	var s entity.Record
	var ss []entity.Record

	ss, err := r.Get()
	if err == nil {
		return ss, nil
	}

	log.Println("Иду в базу")

	q := "select * from session"

	result, err := r.db.Query(q)
	if err != nil {
		return []entity.Record{}, err
	}

	for result.Next() {
		err = result.Scan(&s.ID, &s.Name, &s.Date, &s.Description)
		if err != nil {
			return []entity.Record{}, err
		}

		ss = append(ss, s)
	}

	err = r.Set(ss)
	if err != nil {
		log.Println("cache set error")
	}

	return ss, nil
}

func (r *Repository) Insert(record entity.Record) error {
	r.c.Del(context.Background(), "asd321")

	q := "insert into session(id, name, date, description) values ($1, $2, $3, $4)"

	_, err := r.db.Exec(q, record.ID, record.Name, record.Date, record.Description)
	if err != nil {
		return err
	}

	return nil
}
