package ports

import (
	"context"
	"pgad/internal/core/domain"
)

// Inbound Port (use case API)
type TelemetryUseCase interface {
	FetchLatest(ctx context.Context, limit int) ([]domain.Telemetry, error)
}

// Outbound Port (คลาวด์/แหล่งข้อมูล)
type TelemetrySource interface {
	FetchLatest(ctx context.Context, limit int) ([]domain.Telemetry, error)
}
