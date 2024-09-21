package main

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"orchestrator/router"
	"orchestrator/struct_type"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:      true,
		ServerHeader: "Fiber",
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			return ctx.Status(code).JSON(&struct_type.ErrorResponse{
				Message: "Something went wrong, please try again later.",
				Error:   err.Error(),
			})
		},
	})

	app.Use(recover.New())

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} [${latency}] ${status} - ${method} ${path}\n",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Arjuna Orchestrator 🔥")
	})

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
