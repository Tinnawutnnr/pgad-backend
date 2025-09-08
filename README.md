# PGAD - Telemetry Data Gateway

A Go-based HTTP service that fetches and logs telemetry data (accelerometer readings) from IoT devices. This project follows Clean Architecture principles with a hexagonal architecture pattern.

## 🏗️ Architecture

This project implements **Clean Architecture** with the following layers:

- **Domain Layer**: Core business entities and rules
- **Use Case Layer**: Application business rules and orchestration
- **Adapter Layer**: External interfaces (HTTP handlers, data sources)
- **Infrastructure Layer**: Framework and external dependencies

## 📁 Project Structure

```
PGAD/
├── cmd/
│   └── server/
│       └── main.go                    # Application entry point
├── internal/
│   ├── app/
│   │   └── bootstrap.go               # Dependency injection & app setup
│   ├── core/                         # Core business logic (Clean Architecture)
│   │   ├── domain/
│   │   │   └── telemetry.go          # Domain entities
│   │   ├── ports/
│   │   │   └── telemetry_ports.go    # Interfaces (inbound/outbound ports)
│   │   └── usecase/
│   │       └── telemetry_service.go  # Business logic implementation
│   └── adapters/                     # External adapters
│       ├── inbound/
│       │   └── http/
│       │       └── telemetry_handler.go  # HTTP REST API handler
│       └── outbound/
│           ├── cloud/
│           │   └── http_source.go    # Real cloud/API data source
│           └── mock/
│               └── mock_source.go    # Mock data source for testing
├── go.mod                           # Go module dependencies
└── README.md                        # This file
```

## 🚀 Features

- **Telemetry Data Fetching**: Retrieve accelerometer data (AX, AY, AZ) from IoT devices
- **Dual Data Sources**:
  - Real cloud/API integration (when `CLOUD_BASE_URL` is set)
  - Mock data source for development/testing
- **REST API**: Simple HTTP endpoint to fetch and log telemetry data
- **Clean Architecture**: Testable, maintainable, and loosely coupled design
- **Configurable Limits**: Fetch limits with sensible defaults (1-1000 records)

## 🔧 Installation & Setup

### Prerequisites

- Go 1.25.0 or higher
- Git

### Clone the Repository

```bash
git clone <repository-url>
cd PGAD
```

### Install Dependencies

```bash
go mod tidy
```

## 🏃‍♂️ Running the Application

### Method 1: Using Go Run

```bash
# Run with mock data source (default)
go run cmd/server/main.go

#or

Make run

# Run with real cloud data source
CLOUD_BASE_URL="https://your-api.com" CLOUD_API_KEY="your-api-key" go run cmd/server/main.go
```

### Method 2: Build and Run

```bash
# Build the application
go build -o pgad cmd/server/main.go

#or

Make build

# Run with mock data
./pgad

# Run with cloud data source
CLOUD_BASE_URL="https://your-api.com" CLOUD_API_KEY="your-api-key" ./pgad
```

The server will start on port `:8080` by default.

## 📡 API Endpoints

### GET /api/pull

Fetches the latest telemetry data and logs it to stdout.

**Query Parameters:**

- `limit` (optional): Number of records to fetch (default: 10, max: 1000)

**Example Request:**

```bash
curl "http://localhost:8080/api/pull?limit=5"
```

**Example Response:**

```json
{
  "fetched": 3
}
```

**Console Output:**

```
TEL | dev=esp8266-01 ts=2025-09-08T10:30:15Z ax=0.010 ay=0.020 az=0.980
TEL | dev=esp8266-01 ts=2025-09-08T10:30:16Z ax=0.020 ay=0.010 az=0.990
TEL | dev=esp8266-01 ts=2025-09-08T10:30:17Z ax=-0.010 ay=0.000 az=1.010
```

## ⚙️ Configuration

The application uses environment variables for configuration:

| Variable         | Description                              | Default               |
| ---------------- | ---------------------------------------- | --------------------- |
| `CLOUD_BASE_URL` | Base URL for the cloud telemetry API     | None (uses mock data) |
| `CLOUD_API_KEY`  | API key for cloud service authentication | None                  |

### Example Configuration

```bash
export CLOUD_BASE_URL="https://api.telemetry-service.com"
export CLOUD_API_KEY="your-secret-api-key"
```

## 🧪 Development

### Mock Data Source

When no `CLOUD_BASE_URL` is provided, the application uses a mock data source that generates sample telemetry data:

```go
{DeviceID: "esp8266-01", Timestamp: now, AX: 0.01, AY: 0.02, AZ: 0.98}
```

### Adding New Data Sources

To add a new data source:

1. Implement the `TelemetrySource` interface in `internal/core/ports/telemetry_ports.go`
2. Create your adapter in `internal/adapters/outbound/`
3. Wire it up in `internal/app/bootstrap.go`

### Testing

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests for a specific package
go test ./internal/core/usecase/
```

## 📋 Data Model

### Telemetry Entity

```go
type Telemetry struct {
    DeviceID  string    `json:"deviceId"`  // Device identifier (e.g., "esp8266-01")
    Timestamp time.Time `json:"timestamp"` // UTC timestamp
    AX        float64   `json:"ax"`        // X-axis acceleration
    AY        float64   `json:"ay"`        # Y-axis acceleration
    AZ        float64   `json:"az"`        # Z-axis acceleration
}
```

## 🛠️ Technology Stack

- **Language**: Go 1.25.0
- **Web Framework**: Fiber v2
- **Architecture**: Clean Architecture / Hexagonal Architecture
- **HTTP Client**: Native Go `net/http`
- **Logging**: Standard Go `log` package

## 🏗️ Architecture Principles

1. **Dependency Inversion**: Core business logic depends on abstractions, not concretions
2. **Single Responsibility**: Each component has a single, well-defined purpose
3. **Open/Closed**: Easy to extend with new data sources without modifying existing code
4. **Interface Segregation**: Small, focused interfaces
5. **Separation of Concerns**: Clear boundaries between layers

## 📝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🤝 Support

For support and questions, please open an issue in the GitHub repository.
