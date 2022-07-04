package service

import (
	"context"
	"nails/entity"
	"time"
)

type Repository interface {
	Insert(ctx context.Context, record entity.Record) (entity.Record, error)
	Records(ctx context.Context) ([]entity.Record, error)
}

type UseCase interface {
	RecordAdd(ctx context.Context, ss entity.Record) (entity.Record, error)
	Records(ctx context.Context) ([]entity.Record, error)
}

type service struct {
	repo Repository
}

func New(r Repository) UseCase {
	return &service{repo: r}
}

func (s *service) RecordAdd(ctx context.Context, ss entity.Record) (entity.Record, error) {
	ss.CreatedAt = time.Now()
	ss.Status = entity.StatusActive

	ss, err := s.repo.Insert(ctx, ss)
	if err != nil {
		return entity.Record{}, err
	}

	return ss, nil
}

func (s *service) Records(ctx context.Context) ([]entity.Record, error) {
	ss, err := s.repo.Records(ctx)
	if err != nil {
		return ss, err
	}

	return ss, nil
}
