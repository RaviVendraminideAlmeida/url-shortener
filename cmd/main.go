package main

import (
	"github.com/RaviVendraminideAlmeida/url-shortener/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return handlers.GetURL(c)
	})

	err := app.Listen(":3000")

	if err != nil {
		panic(err)
	}
}
