package domain

import "time"

type Telemetry struct {
	DeviceID  string    `json:"deviceId"`
	Timestamp time.Time `json:"timestamp"`
	AX        float64   `json:"ax"`
	AY        float64   `json:"ay"`
	AZ        float64   `json:"az"`
}
