package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	errorCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "error_counter",
			Help: "Number of errors",
		},
		[]string{"error_type", "context"},
	)
)

type ErrorMetric struct {
	ErrorType    string `json:"error_type"`
	ErrorMessage string `json:"error_message"`
	Context      string `json:"context"`
}

func HandleErrorMetric(c *fiber.Ctx) error {
	var metric ErrorMetric
	if err := c.BodyParser(&metric); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid metric data",
		})
	}

	// Incrémente le compteur d'erreurs avec les labels appropriés
	errorCounter.WithLabelValues(metric.ErrorType, metric.Context).Inc()

	return c.SendStatus(fiber.StatusOK)
}
