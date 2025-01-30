package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Nombre total de requêtes HTTP",
		},
		[]string{"method", "path", "status"},
	)

	errorCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "app_errors_total",
			Help: "Nombre total d'erreurs par type",
		},
		[]string{"error_type", "endpoint"},
	)
)

func PrometheusMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()

		status := fiber.StatusInternalServerError
		if err == nil {
			status = c.Response().StatusCode()
		}

		// Ajoute des logs pour le debug
		log.Printf("Request: %s %s, Status: %d", c.Method(), c.Path(), status)

		if err != nil {
			log.Printf("Error: %v", err)
			errorCounter.WithLabelValues(err.Error(), c.Path()).Inc()
		}

		httpRequestsTotal.WithLabelValues(
			c.Method(),
			c.Path(),
			string(rune(status)),
		).Inc()

		// Si une erreur survient, incrémente le compteur d'erreurs
		if err != nil {
			errorCounter.WithLabelValues(err.Error(), c.Path()).Inc()
		}

		return err
	}
}
