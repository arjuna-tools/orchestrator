package pods

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	v2 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"orchestrator/config"
	"orchestrator/struct_type"
)

// DeletePod [DELETE/:namespace/:podName]
func DeletePod(c *fiber.Ctx) error {
	nameSpace := c.Params("namespace")
	podName := c.Params("podName")
	client := config.ClientSetCoreV1()

	err := client.Pods(nameSpace).Delete(context.TODO(), podName, v2.DeleteOptions{})
	if err != nil {
		return c.Status(500).JSON(&struct_type.ErrorResponse{
			Message: "Unable to deleted pod",
			Error:   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Pod %s deleted successfully", podName),
	})
}
