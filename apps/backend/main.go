package main

import (
	"backend/api/handler"
	"backend/database"
	"backend/repository"
	"backend/service"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	db := database.ConnectDB()
	defer db.Close()

	linkRepo := repository.NewLinkRepository(db)
	linkService := service.NewLinkService(linkRepo)
	linkHandler := handler.NewLinkHandler(linkService)

	app := fiber.New(fiber.Config{
		EnablePrintRoutes: false,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	app.Use(recover.New())

	api := app.Group("/api/v1")
	api.Post("/shorten", linkHandler.CreateShortURL)
	app.Get("/:shortURL", linkHandler.RedirectToURL)

	log.Fatal(app.Listen(":3000"))
}
