package main

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"orchestrator/config"
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

	app.Use(func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		// Check if the Authorization header is missing
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Missing Authorization header",
			})
		}

		if authHeader != config.EnvConfig("API_TOKEN") {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid Authorization Token!",
			})
		}

		// Continue processing the request if the token is valid
		return c.Next()
	})

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
