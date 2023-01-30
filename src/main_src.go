package src

import (
	"dhikrama.com/web/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func MainSrc() {
	engine := html.New("./public/views", ".html")
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})
	routes.RoutesHtml(app)

	app.Listen(":8080")
}
