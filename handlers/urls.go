package handlers

import (
	"github.com/RaviVendraminideAlmeida/url-shortener/database"
	"github.com/RaviVendraminideAlmeida/url-shortener/models"
	"github.com/gofiber/fiber/v2"
)

func GetURLS(c *fiber.Ctx) error {

	urls := []models.URL{}

	database.DB.Db.Find(&urls)

	return c.Status(200).JSON(urls)
}

func ShortenURL(c *fiber.Ctx) error {

	return nil
}
