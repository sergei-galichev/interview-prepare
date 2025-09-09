package main

import (
	"context"
	"github.com/google/uuid"
)

type Service struct {
	works []string
}

func NewService() *Service {
	return &Service{
		works: make([]string, 0),
	}
}

func (s *Service) CreateWork(ctx context.Context) (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	work := "work-" + id.String()

	s.works = append(s.works, work)

	return work, nil
}
