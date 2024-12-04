package routes

import (
	"mainbe/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupTaskRoutes(app *fiber.App) {
	// Customer
	app.Post("/register", controller.RegisterCustomer)

	// Seller
	app.Post("/register/seller", controller.RegisterSeller)

	// Login
	app.Post("/login", controller.Login)

	// My Profile
	app.Get("/profile", controller.GetMyProfile)

}