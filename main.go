package main

import (
	"mainbe/config"
	"mainbe/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Create new fiber app
	app := fiber.New()

	// Connect to database
	db := config.GetDB()

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowHeaders: "*",
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	app.Use(logger.New(logger.Config{
		Format: "${status} - ${method} ${path}\n",
	}))

	// Save database connections in fiber app context
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	// Routes
	routes.SetupTaskRoutes(app)

	// Listen to port 3000
	app.Listen(":3000")
}
