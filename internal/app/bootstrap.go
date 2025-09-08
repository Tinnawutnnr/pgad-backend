package app

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	httpin "pgad/internal/adapters/inbound/http"
	"pgad/internal/adapters/outbound/cloud"
	"pgad/internal/adapters/outbound/mock"
	"pgad/internal/core/ports"
	"pgad/internal/core/usecase"
)

type App struct{ Fiber *fiber.App }

func New() *App {
	f := fiber.New()

	// เลือก Source: ถ้าตั้ง CLOUD_BASE_URL -> ใช้ของจริง, ไม่งั้นใช้ mock
	var src ports.TelemetrySource
	base := os.Getenv("CLOUD_BASE_URL")
	api := os.Getenv("CLOUD_API_KEY")
	if base != "" {
		log.Println("Using CLOUD HTTP source:", base)
		src = cloud.NewHTTPSource(base, api)
	} else {
		log.Println("Using MOCK source (no CLOUD_BASE_URL)")
		src = mock.NewMockSource()
	}

	uc := usecase.NewTelemetryService(src)
	h := httpin.NewTelemetryHandler(uc)

	apiGroup := f.Group("/api")
	h.Register(apiGroup)

	return &App{Fiber: f}
}

func (a *App) Start(addr string) error { return a.Fiber.Listen(addr) }
