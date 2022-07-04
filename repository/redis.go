package repository

import (
	"context"
	"encoding/json"
)

var s Records

func (r *Repository) Set(ctx context.Context, record Records) error {
	j, err := json.Marshal(record)
	if err != nil {
		return err
	}

	err = r.c.Set(ctx, "asd321", j, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Get(ctx context.Context) (Records, error) {
	err := r.c.Get(ctx, "asd321").Scan(&s)
	if err != nil {
		return nil, err
	}

	return s, nil
}
