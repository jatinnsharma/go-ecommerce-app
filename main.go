package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("I am main function")

	// Create a new Fiber instance
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Start the server on localhost:9000
	if err := app.Listen("localhost:9000"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
