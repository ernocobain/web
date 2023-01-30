package handlers

import "github.com/gofiber/fiber/v2"

// Handlers home started code

func HandlerHome(h *fiber.Ctx) error {
	return h.Render("pages/home/index", fiber.Map{
		"title": "Hello world",
	})
}

// Handlers about started code

func HandlerAbout(h *fiber.Ctx) error {
	return h.JSON(fiber.Map{
		"welcome": "Hello World",
	})
}

// Handlers contact started code

func HandlerContact(h *fiber.Ctx) error {
	return h.JSON(fiber.Map{
		"welcome": "Hello World",
	})
}
