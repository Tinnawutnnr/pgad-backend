package httpin

import (
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"pgad/internal/core/ports"
)

type TelemetryHandler struct {
	uc ports.TelemetryUseCase
}

func NewTelemetryHandler(uc ports.TelemetryUseCase) *TelemetryHandler {
	return &TelemetryHandler{uc: uc}
}

func (h *TelemetryHandler) Register(r fiber.Router) {
	r.Get("/pull", h.pullAndLog)
}

func (h *TelemetryHandler) pullAndLog(c *fiber.Ctx) error {
	limitStr := c.Query("limit", "10")
	limit, _ := strconv.Atoi(limitStr)

	ctx := c.Context() // Fiber's fasthttp ctx; we'll ignore and use Background
	_ = ctx

	data, err := h.uc.FetchLatest(c.Context(), limit)
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}

	// log แบบง่าย ๆ ไป stdout (ดูใน terminal)
	for _, t := range data {
		log.Printf("TEL | dev=%s ts=%s ax=%.3f ay=%.3f az=%.3f",
			t.DeviceID, t.Timestamp.Format(time.RFC3339), t.AX, t.AY, t.AZ)
	}

	// ตอบกลับแบบเบา ๆ ว่าดึงได้กี่เรคคอร์ด
	return c.JSON(fiber.Map{"fetched": len(data)})
}
