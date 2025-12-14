package services

import (
	models "go-boilerplate/internal/health/model"
	repositories "go-boilerplate/internal/health/repository"
)

type HealthService struct {
	repo *repositories.HealthRepository
}

func NewHealthService(repo *repositories.HealthRepository) *HealthService {
	return &HealthService{repo: repo}
}

func (s *HealthService) CheckHealth() models.HealthResponse {
	status := "ok"
	if err := s.repo.PingDB(); err != nil {
		status = err.Error()
	}

	return models.HealthResponse{
		Status:  "ok",
		DB:      status,
		Version: "1.0.0",
	}
}
