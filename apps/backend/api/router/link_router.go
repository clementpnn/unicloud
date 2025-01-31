package router

import (
	"backend/api/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupLinkRoutes(app *fiber.App, linkHandler *handler.LinkHandler) {
	api := app.Group("/api/v1")
	api.Post("/shorten", linkHandler.CreateShortURL)
	api.Get("/:shortURL", linkHandler.RedirectToURL)
}
