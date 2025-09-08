package cloud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"pgad/internal/core/domain"
	"pgad/internal/core/ports"
)

type HTTPSource struct {
	BaseURL string
	APIKey  string
	Client  *http.Client
}

func NewHTTPSource(baseURL, apiKey string) ports.TelemetrySource {
	return &HTTPSource{
		BaseURL: baseURL,
		APIKey:  apiKey,
		Client:  &http.Client{Timeout: 10 * time.Second},
	}
}

func (s *HTTPSource) FetchLatest(ctx context.Context, limit int) ([]domain.Telemetry, error) {
	u, err := url.Parse(s.BaseURL)
	if err != nil {
		return nil, err
	}
	// สมมติ endpoint: GET /telemetry?limit=N  (ปรับตามของจริงได้)
	u.Path = "/telemetry"
	q := u.Query()
	q.Set("limit", fmt.Sprintf("%d", limit))
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	if s.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+s.APIKey)
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("cloud http status %d", resp.StatusCode)
	}

	var out []domain.Telemetry
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}
