package repository

import (
	// "github.com/dgrijalva/jwt-go"
	"mainbe/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	// "time"
)

func CreateCustomer(db *gorm.DB, user *model.User) error {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Jika id role tidak diisi, maka set default role ke 3
	if user.IdRole == 0 {
		user.IdRole = 3
	}

	// Create user
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func CreateSeller(db *gorm.DB, user *model.User) error {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Jika id role tidak diisi, maka set default role ke 2
	if user.IdRole == 0 {
		user.IdRole = 2
	}

	// Create user
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func GenerateToken(user *model.User) (string, error) {
	// Set claims
	claims := &model.JWTClaims{
		IdUser: user.IdUser,
		IdRole: user.IdRole,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret_key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetUserByUsername(db *gorm.DB, username string) (model.User, error) {
	var user model.User

	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func GetUserById(db *gorm.DB, idUser int) (model.User, error) {
	var user model.User

	if err := db.First(&user, idUser).Error; err != nil {
		return user, err
	}

	return user, nil
}

func GetUserByEmail(db *gorm.DB, email string) (model.User, error) {
	var user model.User

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func UpdateUser(db *gorm.DB, user *model.User) error {
	if err := db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}