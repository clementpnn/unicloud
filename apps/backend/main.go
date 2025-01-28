package main

import (
	"backend/api/handler"
	"backend/api/router"
	"backend/database"
	"backend/repository"
	"backend/service"

	_ "backend/api/doc"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

var (
	crawlerHandler *handler.CrawlerHandler
	searchHandler  *handler.SearchHandler
)

func init() {
	db := database.ConnectDB()

	domainRepo := repository.NewDomainRepository(db)
	pageRepo := repository.NewPageRepository(db)
	wordRepo := repository.NewWordRepository(db)
	searchRepo := repository.NewSearchRepository(db)

	domainService := service.NewDomainService(domainRepo)
	pageService := service.NewPageService(pageRepo)
	wordService := service.NewWordService(wordRepo, pageRepo)
	searchService := service.NewSearchService(searchRepo)

	crawlerHandler = handler.NewCrawlerHandler(domainService, pageService, wordService)
	searchHandler = handler.NewSearchHandler(searchService)

}

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(recover.New())

	app.Get("/swagger/*", swagger.HandlerDefault)
	router.CrawlerRoutes(app, crawlerHandler)
	router.SearchRoutes(app, searchHandler)
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	app.Listen(":3000")
}
