package main

import "github.com/gofiber/fiber/v2"

func main() {
	web := fiber.New()

	web.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello World",
		})
	})

	web.Listen(":8080")
}
