package main

import (
	"docker-go/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.SetupDatabase()
	config.SetupRoutes(app)

	app.Listen(":3000")
}
