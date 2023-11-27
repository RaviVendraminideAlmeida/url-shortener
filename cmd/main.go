package main

import (
	"log"

	"github.com/RaviVendraminideAlmeida/url-shortener/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	database.ConnectDB()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "URL Shortener",
		})
	})

	log.Fatal(app.Listen(":3000"))

}
