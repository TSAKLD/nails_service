package repository

import (
	"context"
	"log"
	"nails/entity"
)

func (r *Repository) Records(ctx context.Context) ([]entity.Record, error) {
	ss, err := r.Get(ctx)
	if err == nil {
		return ss, nil
	}

	q := "select * from session"

	result, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return []entity.Record{}, err
	}

	for result.Next() {
		var s entity.Record

		err = result.Scan(&s.ID, &s.Name, &s.Date, &s.Description)
		if err != nil {
			return []entity.Record{}, err
		}

		ss = append(ss, s)
	}

	err = result.Close()
	if err != nil {
		log.Println("Unable to close request")
	}

	err = r.Set(ctx, ss)
	if err != nil {
		log.Println("cache set error")
	}

	return ss, nil
}

func (r *Repository) Insert(ctx context.Context, record entity.Record) (entity.Record, error) {
	q := "INSERT into session(name, date, created_at, description, status)" +
		" VALUES ($1, $2, $3, $4, $5) RETURNING *"

	row := r.db.QueryRowContext(ctx, q, record.Name,
		record.Date, record.CreatedAt, record.Description, record.Status)

	err := row.Scan(&record.ID, &record.Name, &record.Date, &record.Description, &record.CreatedAt, &record.Status)
	if err != nil {
		return record, err
	}

	err = r.c.Del(ctx, "asd321").Err()
	if err != nil {
		log.Println("cache clear error")
	}
	return record, nil
}
