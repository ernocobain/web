package routes

import (
	"dhikrama.com/web/src/handlers"
	"github.com/gofiber/fiber/v2"
)

func RoutesHtml(app *fiber.App) {
	app.Get("", handlers.HandlerHome)
	app.Get("about", handlers.HandlerAbout)
	app.Get("contact", handlers.HandlerContact)
}
