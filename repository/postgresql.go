package repository

import (
	"context"
	"log"
	"nails/entity"
)

func (r *Repository) Records() ([]entity.Record, error) {
	ss, err := r.Get()
	if err == nil {
		return ss, nil
	}

	q := "select * from session"

	result, err := r.db.Query(q)
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

	err = r.Set(ss)
	if err != nil {
		log.Println("cache set error")
	}

	return ss, nil
}

func (r *Repository) Insert(record entity.Record) (entity.Record, error) {
	q := "INSERT into session(name, date, description)" +
		" VALUES ($1, $2, $3) RETURNING *"

	row := r.db.QueryRowContext(context.Background(), q, record.Name,
		record.Date, record.Description)

	err := row.Scan(&record.ID, &record.Name, &record.Date, &record.Description)
	if err != nil {
		return record, err
	}

	r.c.Del(context.Background(), "asd321")

	return record, nil
}
