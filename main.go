package main

import (
	"github.com/gofiber/fiber/v3/middleware/logger"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v3"
)

func main() {
	log.Println("Starting application...")

	client := &http.Client{}
	authHeader := "Bearer " + os.Getenv("openApiKey")
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/joke", func(c fiber.Ctx) error {
		return jokeHandler(c, client, &authHeader)
	})

	app.Get("/image", func(c fiber.Ctx) error {
		description := c.Query("description")
		if description == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Description query parameter is required")
		}
		return imageHandler(c, client, &description, &authHeader)
	})

	app.Get("/audio", func(c fiber.Ctx) error {
		text := c.Query("text")
		if text == "" {
			return c.Status(fiber.StatusBadRequest).SendString("Text query parameter is required")
		}
		return audioHandler(c, client, &text, &authHeader)
	})

	log.Fatal(app.Listen(":3000", fiber.ListenConfig{DisableStartupMessage: true}))
}
