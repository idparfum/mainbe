package controller

import (
	"mainbe/model"
	repo "mainbe/repository"

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

	// Validasi id user (id user akan digenerate, jadi pastikan kosong)
	if user.IdUser != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Id user tidak boleh diisi, id akan digenerate otomatis",
		})
	}

	// Menyimpan data user ke database
	if err := repo.CreateCustomer(db, &user); err != nil {
		if err.Error() == "invalid phone number format" || err.Error() == "invalid email format" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
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

	// Validasi id user (id user akan digenerate, jadi pastikan kosong)
	if user.IdUser != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Id user tidak boleh diisi, id akan digenerate otomatis",
		})
	}

	// Menyimpan data user ke database
	if err := repo.CreateSeller(db, &user); err != nil {
		if err.Error() == "invalid phone number format" || err.Error() == "invalid email format" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
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
			"message": "Username salah",
		})
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Password salah",
		})
	}

	// Generate token with the user ID from the database (userData.ID)
	token, err := repo.GenerateToken(userData.IdUser)  // Pass the actual user ID from the DB
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Generate token failed",
		})
	}

	// Return the token in the response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}



func GetMyProfile(c *fiber.Ctx) error {
	// Ambil claims dari Locals
	claims := c.Locals("claims").(*model.JWTClaims)

	// Ambil user ID dari claims
	userId := claims.IdUser

	// Ambil database instance dari Locals
	db := c.Locals("db").(*gorm.DB)

	// Cari data user berdasarkan ID
	user, err := repo.GetUserById(db, userId)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "User tidak ditemukan")
	}

	return c.JSON(fiber.Map{
		"user": user,
	})
}
