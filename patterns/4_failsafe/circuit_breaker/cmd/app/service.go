package main

import (
	"context"
	"errors"
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
	if len(s.works) >= 5 {
		return "full queue", errors.New("failed create work: full queue")
	}

	id, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	work := "work-" + id.String()

	s.works = append(s.works, work)

	return work, nil
}
