package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"orchestrator/router"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:      true,
		ServerHeader: "Fiber",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Arjuna Orchestrator 🔥")
	})

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
