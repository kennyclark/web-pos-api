package server

import (
	"github.com/goccy/go-json"
	"github.com/kennyclark/web-pos-api/internal/routes"
	"github.com/kennyclark/web-pos-api/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func New() (*fiber.App, error) {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(recover.New())
	app.Use(logger.New())

	// CORS middleware
	allowedOrigins := "*"
	mode := utils.GetEnvOrDefault("MODE", "development")
	if mode == "production" {
		allowedOrigins = utils.GetEnvOrDefault("CORS_ALLOWED_ORIGINS", "")
	}
	app.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
		AllowHeaders: "Origin, Content-Type, Accept-Encoding, Authorization",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	// health check endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Web-POS API.")
	})

	// API routes
	api := app.Group("/api")
	routes.Categories(api)

	return app, nil
}
