package routes

import (
	"dhikrama.com/web/src/handlers"
	"github.com/gofiber/fiber/v2"
)

func RoutersApi(app *fiber.App) {
	app.Get("/api", handlers.UserGetAll)
}
