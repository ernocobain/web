package src

import (
	"dhikrama.com/web/src/routes"
	"github.com/gofiber/fiber/v2"
)

func MainSrc() {
	app := fiber.New()

	routes.RoutesHtml(app)

	app.Listen(":8080")
}
