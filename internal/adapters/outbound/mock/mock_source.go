package mock

import (
	"context"
	"time"

	"pgad/internal/core/domain"
	"pgad/internal/core/ports"
)

type MockSource struct{}

func NewMockSource() ports.TelemetrySource { return &MockSource{} }

func (m *MockSource) FetchLatest(ctx context.Context, limit int) ([]domain.Telemetry, error) {
	if limit <= 0 { limit = 3 }
	now := time.Now().UTC()
	data := []domain.Telemetry{
		{DeviceID: "esp8266-01", Timestamp: now.Add(-3 * time.Second), AX: 0.01, AY: 0.02, AZ: 0.98},
		{DeviceID: "esp8266-01", Timestamp: now.Add(-2 * time.Second), AX: 0.02, AY: 0.01, AZ: 0.99},
		{DeviceID: "esp8266-01", Timestamp: now.Add(-1 * time.Second), AX: -0.01, AY: 0.00, AZ: 1.01},
	}
	if limit < len(data) { data = data[:limit] }
	return data, nil
}
