package service

import (
	"context"
)

type HealthService struct {
	postgresPing func(ctx context.Context) error
	redisPing    func(ctx context.Context) error
}

func NewHealthService(
	postgresPing func(ctx context.Context) error,
	redisPing func(ctx context.Context) error,
) *HealthService {
	return &HealthService{
		postgresPing: postgresPing,
		redisPing:    redisPing,
	}
}

func (s *HealthService) Check(ctx context.Context) (map[string]string, error) {
	if err := s.postgresPing(ctx); err != nil {
		return map[string]string{
			"status": "postgres unavailable",
		}, err
	}

	if err := s.redisPing(ctx); err != nil {
		return map[string]string{
			"status": "redis unavailable",
		}, err
	}

	return map[string]string{
		"status": "ok",
	}, nil
}
