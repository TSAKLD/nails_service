package repository

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v9"
	"nails/entity"
)

type Records []entity.Record

var s Records

func (r *Records) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &r)
}

func (r *Records) MarshalBinary() (data []byte, err error) {
	return json.Marshal(data)
}

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return client
}

func (r *Repository) Set(record Records) error {
	j, err := json.Marshal(record)
	if err != nil {
		return err
	}

	err = r.c.Set(context.Background(), "asd321", j, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Get() (Records, error) {
	err := r.c.Get(context.Background(), "asd321").Scan(&s)
	if err != nil {
		return nil, err
	}

	return s, nil
}
