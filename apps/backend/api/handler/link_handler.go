package handler

import (
	"backend/service"

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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	link, err := h.linkService.CreateShortURL(req.URL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(link)
}

func (h *LinkHandler) RedirectToURL(c *fiber.Ctx) error {
	shortURL := c.Params("shortURL")

	originalURL, err := h.linkService.GetOriginalURL(shortURL)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "URL not found",
		})
	}


	return c.Redirect(originalURL, fiber.StatusTemporaryRedirect)
}

type CreateLinkRequest struct {
	URL string `json:"url"`
}
