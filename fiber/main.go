package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Get("/hi/:name", func(c *fiber.Ctx) error {
		return c.SendString("HI " + c.Params("name"))
	})
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(":3000"))
}
