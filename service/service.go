package service

import "nails/entity"

type Repository interface {
	Insert(record entity.Record) (entity.Record, error)
	Records() ([]entity.Record, error)
}

type UseCase interface {
	RecordAdd(ss entity.Record) (entity.Record, error)
	Records() ([]entity.Record, error)
}

type service struct {
	repo Repository
}

func New(r Repository) UseCase {
	return &service{repo: r}
}

func (s *service) RecordAdd(ss entity.Record) (entity.Record, error) {
	ss, err := s.repo.Insert(ss)
	if err != nil {
		return entity.Record{}, err
	}

	return ss, nil
}

func (s *service) Records() ([]entity.Record, error) {
	ss, err := s.repo.Records()
	if err != nil {
		return ss, err
	}

	return ss, nil
}
