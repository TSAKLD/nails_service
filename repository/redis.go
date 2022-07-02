package repository

import (
	"context"
	"encoding/json"
)

var s Records

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
