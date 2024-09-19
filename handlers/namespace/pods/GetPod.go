package pods

import (
	"context"
	"github.com/gofiber/fiber/v2"
	v2 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"orchestrator/config"
	"orchestrator/struct_type"
)

// GetPod [GET/:namespace/:podName]
func GetPod(c *fiber.Ctx) error {
	nameSpace := c.Params("namespace")
	podName := c.Params("podName")
	client := config.ClientSetCoreV1()

	pod, err := client.Pods(nameSpace).Get(context.TODO(), podName, v1.GetOptions{})
	if err != nil {
		return c.Status(500).JSON(&struct_type.ErrorResponse{
			Message: "Unable to fetch pod info",
			Error:   err.Error(),
		})
	}

	return c.JSON(&struct {
		Name       string         `json:"name"`
		UID        string         `json:"uid"`
		Node       string         `json:"node"`
		PodIP      string         `json:"podIP"`
		Status     string         `json:"status"`
		StartTime  string         `json:"startTime"`
		Containers []v2.Container `json:"containers" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,2,rep,name=containers"`
	}{
		Name:       pod.Name,
		UID:        string(pod.UID),
		Node:       pod.Spec.NodeName,
		PodIP:      pod.Status.PodIP,
		Status:     string(pod.Status.Phase),
		StartTime:  pod.Status.StartTime.String(),
		Containers: pod.Spec.Containers,
	})
}
