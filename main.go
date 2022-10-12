package main

import (
	"docker-go/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.SetupDatabase()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
