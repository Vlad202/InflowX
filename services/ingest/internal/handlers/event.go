package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vlad202/inflow-x/ingest/internal/queue"
)

type Event struct {
	Type  string                 `json:"type"`
	Data  map[string]interface{} `json:"data"`
	Token string                 `json:"token"`
}

func HandleEvent(c *fiber.Ctx) error {
	event := new(Event)

	if err := c.BodyParser(event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	go queue.PushEvent(event)

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": "ok"})
}
