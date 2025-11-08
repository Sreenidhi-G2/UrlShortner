package handler

import (
	"shorten-service/internal/service"

	"github.com/gofiber/fiber/v2"
)

type URLHandler struct {
	service *service.URLService
}

func NewURLHandler(service *service.URLService) *URLHandler {
	return &URLHandler{service}
}

// POST /shorten
func (h *URLHandler) ShortenURL(c *fiber.Ctx) error {
	type request struct {
		URL string `json:"url"`
	}

	var body request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	shortKey, err := h.service.ShortenURL(body.URL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"short_url": "http://localhost:8001/" + shortKey,
	})
}

// GET /:shortKey
func (h *URLHandler) Redirect(c *fiber.Ctx) error {
	shortKey := c.Params("shortKey")

	originalURL, err := h.service.GetOriginalURL(shortKey)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "URL not found"})
	}

	return c.Redirect(originalURL, fiber.StatusTemporaryRedirect)
}
