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

func (r *Repository) Insert(record entity.Record) error {
	r.c.Del(context.Background(), "asd321")

	q := "insert into session(id, name, date, description) values ($1, $2, $3, $4)"

	_, err := r.db.Exec(q, record.ID, record.Name, record.Date, record.Description)
	if err != nil {
		return err
	}

	return nil
}
