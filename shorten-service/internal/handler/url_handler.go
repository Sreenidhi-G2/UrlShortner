package handler

import (
	"shorten-service/internal/service"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type URLHandler struct {
	service  *service.URLService
	BASE_URL string
}

func NewURLHandler(service *service.URLService, BASE_URL string) *URLHandler {
	return &URLHandler{service: service, BASE_URL: BASE_URL}

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

	if strings.HasPrefix(body.URL, h.BASE_URL+"/") {
		shortKey := strings.TrimPrefix(body.URL, h.BASE_URL+"/")

		originalURL, err := h.service.GetOriginalURL(shortKey)
		if err == nil {
			// Return the original long URL instead of creating a new short URL
			return c.JSON(fiber.Map{
				"short_url": originalURL,
			})
		}
	}

	shortKey, err := h.service.ShortenURL(body.URL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"short_url": h.BASE_URL + "/" + shortKey,
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

// GET /resolve/:shortKey
func (h *URLHandler) ResolveURL(c *fiber.Ctx) error {
	shortKey := c.Params("shortKey")

	originalURL, err := h.service.GetOriginalURL(shortKey)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "URL not found"})
	}

	return c.JSON(fiber.Map{
		"original_url": originalURL,
	})
}
