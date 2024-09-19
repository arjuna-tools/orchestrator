package pods

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	v1 "k8s.io/api/core/v1"
	v2 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"orchestrator/config"
	"orchestrator/struct_type"
)

type PodBodyOptions struct {
	Labels     map[string]string
	Containers []v1.Container
}

// CreatePod [POST/:namespace/:podName]
func CreatePod(c *fiber.Ctx) error {
	nameSpace := c.Params("namespace")
	podName := c.Params("podName")
	client := config.ClientSetCoreV1()

	body := new(PodBodyOptions)
	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(&struct_type.ErrorResponse{
			Message: "Unable to parse request body",
			Error:   err.Error(),
		})
	}

	pod := &v1.Pod{
		ObjectMeta: v2.ObjectMeta{
			Name:   podName,
			Labels: body.Labels,
		},
		Spec: v1.PodSpec{
			Containers: body.Containers,
		},
	}

	createdPod, err := client.Pods(nameSpace).Create(context.TODO(), pod, v2.CreateOptions{})
	if err != nil {
		return c.Status(500).JSON(&struct_type.ErrorResponse{
			Message: "Unable to create pod",
			Error:   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Pod %s created successfully", createdPod.Name),
	})
}
