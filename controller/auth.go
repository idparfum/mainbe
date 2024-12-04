package controller

import (
	"mainbe/model"
	repo "mainbe/repository"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterCustomer(c *fiber.Ctx) error {
	var user model.User

	db := c.Locals("db").(*gorm.DB)

	// Parse body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request body invalid",
		})
	}

	// Menyimpan data user ke database
	if err := repo.CreateCustomer(db, &user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Registrasi Gagal",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Registrasi Berhasil!",
	})
}

func RegisterSeller(c *fiber.Ctx) error {
	var user model.User

	db := c.Locals("db").(*gorm.DB)

	// Parse body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request body invalid",
		})
	}

	// Menyimpan data user ke database
	if err := repo.CreateSeller(db, &user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Registrasi Gagal",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Registrasi Berhasil!",
	})
}

func Login(c *fiber.Ctx) error {
	var user model.User

	db := c.Locals("db").(*gorm.DB)

	// Parse body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request body invalid",
		})
	}

	// Get user by username
	userData, err := repo.GetUserByUsername(db, user.Nama)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Username atau password salah",
		})
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Username atau password salah",
		})
	}

	// Generate token
	token, err := repo.GenerateToken(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Generate token failed",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func GetMyProfile(c *fiber.Ctx) error {
	// Get token from header
	tokenString := c.Get("login")
	if tokenString == "" {
		return fiber.NewError(fiber.StatusNotFound, "Token not found")
	}

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil
	})
	if err != nil {
		return err
	}

	// Memeriksa apakah token valid
	if !token.Valid {
		return fiber.NewError(fiber.StatusBadRequest, "Token invalid")
	}

	// Ekstrak claims dari token
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["id_user"].(float64)
	db := c.Locals("db").(*gorm.DB)

	// Get user by ID
	userData, err := repo.GetUserById(db, int(userID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"user": userData,
	})
}