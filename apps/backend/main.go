package main

import (
	"backend/api/handler"
	"backend/database"
	"backend/middleware"
	"backend/repository"
	"backend/service"
	"net"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	db := database.ConnectDB()
	defer db.Close()

	app := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: false,
	})

	app.Use(middleware.PrometheusMiddleware())
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "GET,POST,HEAD,OPTIONS,PUT,DELETE,PATCH",
		AllowHeaders:     "*",
		AllowCredentials: true,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

	app.Use(func(c *fiber.Ctx) error {
		return c.Next()
	})

	app.Use(recover.New())

	linkRepo := repository.NewLinkRepository(db)
	linkService := service.NewLinkService(linkRepo)
	linkHandler := handler.NewLinkHandler(linkService)

	api := app.Group("/api/v1")
	api.Post("/shorten", linkHandler.CreateShortURL)
	api.Post("/metrics/error", handler.HandleErrorMetric)
	app.Get("/:shortURL", linkHandler.RedirectToURL)

	port := "3001"
	for i := 0; i < 10; i++ {
		listener, err := net.Listen("tcp", ":"+port)
		if err != nil {
			if strings.Contains(err.Error(), "address already in use") {
				portNum, _ := strconv.Atoi(port)
				port = strconv.Itoa(portNum + 1)
				continue
			}
		}
		listener.Close()
		break
	}

	if err := app.Listen(":" + port); err != nil {
	}
}
