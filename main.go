package main

import "github.com/gofiber/fiber/v2"

func main() {
	// Creating a fiber app
	app := fiber.New()

	// Add routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Start app
	app.Listen(":3000")
}
