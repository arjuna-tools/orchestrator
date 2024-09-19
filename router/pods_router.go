package router

import (
	"github.com/gofiber/fiber/v2"
	"orchestrator/handlers/namespace/pods"
)

func SetupPodsRouter(app *fiber.App) {
	app.Get("/:namespace/:podName", pods.GetPod)
	app.Post("/:namespace/:podName", pods.CreatePod)
	app.Delete("/:namespace/:podName", pods.DeletePod)
}
