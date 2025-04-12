package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/vlad202/inflow-x/ingest/internal/handlers"
)

func main() {
	app := fiber.New()

	app.Post("/events", handlers.HandleEvent)

	log.Fatal(app.Listen(":8080"))
}
