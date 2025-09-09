package main

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

var count int

type Service struct {
	works []string
}

func NewService() *Service {
	return &Service{
		works: make([]string, 0),
	}
}

func (s *Service) CreateWork(ctx context.Context) (string, error) {
	count++

	if count <= 4 {
		return "", errors.New("intention fail")
	} else {
		id, err := uuid.NewV7()
		if err != nil {
			return "", err
		}

		work := "work-" + id.String()

		s.works = append(s.works, work)

		return work, nil
	}
}
