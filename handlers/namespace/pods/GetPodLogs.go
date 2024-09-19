package pods

import (
	"bufio"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	v1 "k8s.io/api/core/v1"
	"orchestrator/config"
	"time"
)

func int64Ptr(i int64) *int64 {
	return &i
}

// GetPodLogs [GET/:namespace/:podName/logs]
func GetPodLogs(c *fiber.Ctx) error {
	nameSpace := c.Params("namespace")
	podName := c.Params("podName")
	client := config.ClientSetCoreV1()

	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("Transfer-Encoding", "chunked")

	c.Status(fiber.StatusOK).Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		podLogOpts := v1.PodLogOptions{
			TailLines:  int64Ptr(100),
			Follow:     true,
			Timestamps: true,
		}

		req := client.Pods(nameSpace).GetLogs(podName, &podLogOpts)
		podLogs, err := req.Stream(context.TODO())
		if err != nil {
			fmt.Fprintf(w, "failed to open log stream: %v\n\n", err)
			w.Flush()
			return
		}

		// 4096 Buffer in each Chunk
		buf := make([]byte, 4096)
		for {
			numBytes, err := podLogs.Read(buf)
			if numBytes > 0 {
				logMessage := string(buf[:numBytes])
				fmt.Fprintf(w, logMessage)

				// Flush the writer to ensure data is sent immediately
				w.Flush()
			}

			// Handle errors or EOF
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Errorf("error reading log stream: %v", err)
			}

			time.Sleep(2 * time.Second)
		}
	})

	return nil
}
