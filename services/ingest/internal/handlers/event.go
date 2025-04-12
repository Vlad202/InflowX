package handlers

import (
	"github.com/vlad202/inflow-x/ingest/internal/queue"

	"github.com/gofiber/fiber/v2"
)

type Event struct {
	Type  string                 `json:"type"`
	Data  map[string]interface{} `json:"data"`
	Token string                 `json:"token"` // если нужен auth
}

func HandleEvent(c *fiber.Ctx) error {
	event := new(Event)

	if err := c.BodyParser(event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	// Здесь можно вставить антиспам / валидацию
	go queue.PushEvent(event) // асинхронно отправим в Redis Streams

	return c.SendStatus(fiber.StatusAccepted)
}
