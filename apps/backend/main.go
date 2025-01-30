package main

import (
	"backend/api/handler"
	"backend/database"
	"backend/middleware"
	"backend/repository"
	"backend/service"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	db := database.ConnectDB()
	defer db.Close()

	app := fiber.New()

	app.Use(middleware.PrometheusMiddleware())
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))
	app.Use(recover.New())

	linkRepo := repository.NewLinkRepository(db)
	linkService := service.NewLinkService(linkRepo)
	linkHandler := handler.NewLinkHandler(linkService)

	api := app.Group("/api/v1")
	api.Post("/shorten", linkHandler.CreateShortURL)
	api.Post("/metrics/error", handler.HandleErrorMetric)
	app.Get("/:shortURL", linkHandler.RedirectToURL)

	log.Fatal(app.Listen(":3000"))
}
