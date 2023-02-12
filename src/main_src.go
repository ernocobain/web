package src

import (
	"crypto/tls"
	"log"

	"dhikrama.com/web/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"golang.org/x/crypto/acme/autocert"
)

func MainSrc() {
	engine := html.New("./public/views", ".html")
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	m := &autocert.Manager{
		Prompt: autocert.AcceptTOS,
		// Replace with your domain
		HostPolicy: autocert.HostWhitelist("13.250.14.237"),
		// Folder to store the certificates
		Cache: autocert.DirCache("./certs"),
	}

	cfg := &tls.Config{
		// Get Certificate from Let's Encrypt
		GetCertificate: m.GetCertificate,
		// By default NextProtos contains the "h2"
		// This has to be removed since Fasthttp does not support HTTP/2
		// Or it will cause a flood of PRI method logs
		// http://webconcepts.info/concepts/http-method/PRI
		NextProtos: []string{
			"http/1.1", "acme-tls/1",
		},
	}
	ln, err := tls.Listen("tcp", ":443", cfg)
	if err != nil {
		panic(err)
	}

	app.Static("/", "./public/static")
	routes.RoutesHtml(app)
	routes.RoutersApi(app)

	log.Fatal(app.Listener(ln))
}
