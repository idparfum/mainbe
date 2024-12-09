package routes

import (
	"mainbe/controller"
	"mainbe/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupTaskRoutes(app *fiber.App) {
	// Customer Routes
	customerRoutes := app.Group("/cust")
	customerRoutes.Post("/register", controller.RegisterCustomer)

	// Seller Routes
	sellerRoutes := app.Group("/seller")
	sellerRoutes.Post("/register", controller.RegisterSeller)

	// Auth Routes (Login doesn't require JWT)
	auth := app.Group("/auth")
	auth.Post("/login", controller.Login)

	// User Routes (requires JWT middleware)
	protected := app.Group("/u")
	protected.Use(middleware.JWTMiddleware("secret"))
	protected.Get("/profile", controller.GetMyProfile)
}
