package router

import (
	"github.com/gofiber/fiber/v2"
	"orchestrator/handlers/namespace/pods"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/:namespace/:podName", pods.GetPod)
}