package repository

import (
	"mainbe/model"
	"mainbe/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateCustomer(db *gorm.DB, user *model.User) error {
	// Generate random ID before creating user
	user.IdUser = utils.GenerateRandomID(1, 10000) // ID acak antara 1 dan 10000

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Validasi phone
	if err := ValidatePhone(user.Phone); err != nil {
		return err
	}

	// Validasi email
	if err := ValidateEmail(user.Email); err != nil {
		return err
	}

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
	// Generate random ID before creating user
	user.IdUser = utils.GenerateRandomID(1, 10000) // ID acak antara 1 dan 10000

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Validasi phone
	if err := ValidatePhone(user.Phone); err != nil {
		return err
	}

	// Validasi email
	if err := ValidateEmail(user.Email); err != nil {
		return err
	}

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

func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"id_user": userID,  // Pastikan ID yang valid digunakan di sini
		"exp":    time.Now().Add(time.Hour * 24).Unix(),  // Token berlaku selama 1 hari
	}

	// Generate token dengan claims yang sudah disiapkan
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}


func GetUserByUsername(db *gorm.DB, nama string) (*model.User, error) {
	var user model.User

	result := db.Where("nama = ?", nama).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func GetUserById(db *gorm.DB, userID uint) (*model.User, error) {
	var user model.User

	result := db.First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
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
