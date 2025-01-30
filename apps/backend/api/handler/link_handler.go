package handler

import (
	"backend/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

type LinkHandler struct {
	linkService *service.LinkService
}

func NewLinkHandler(linkService *service.LinkService) *LinkHandler {
	return &LinkHandler{linkService: linkService}
}

// @Summary Créer une URL courte
// @Description Crée une version courte d'une URL longue
// @Tags links
// @Accept json
// @Produce json
// @Param url body CreateLinkRequest true "URL longue"
// @Success 200 {object} model.Link
// @Router /api/v1/shorten [post]
func (h *LinkHandler) CreateShortURL(c *fiber.Ctx) error {
	var req CreateLinkRequest
	if err := c.BodyParser(&req); err != nil {
		log.Printf("Erreur de parsing: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	link, err := h.linkService.CreateShortURL(req.URL)
	if err != nil {
		log.Printf("Erreur de création: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(link)
}

func (h *LinkHandler) RedirectToURL(c *fiber.Ctx) error {
	shortURL := c.Params("shortURL")

	log.Printf("Trying to redirect shortURL: %s", shortURL)

	link, err := h.linkService.GetByShortURL(shortURL)
	if err != nil {
		log.Printf("Error finding URL: %v", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "URL not found",
		})
	}

	return c.Redirect(link.LongURL, fiber.StatusMovedPermanently)
}

type CreateLinkRequest struct {
	URL string `json:"url"`
}
