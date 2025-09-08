package usecase

import (
	"context"

	"pgad/internal/core/domain"
	"pgad/internal/core/ports"
)

type telemetryService struct {
	src ports.TelemetrySource
}

func NewTelemetryService(src ports.TelemetrySource) ports.TelemetryUseCase {
	return &telemetryService{src: src}
}

func (s *telemetryService) FetchLatest(ctx context.Context, limit int) ([]domain.Telemetry, error) {
	if limit <= 0 || limit > 1000 {
		limit = 100
	}
	return s.src.FetchLatest(ctx, limit)
}
