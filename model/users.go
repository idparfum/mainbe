package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	IdUser    uint      `gorm:"primaryKey;column:id_user" json:"id_user"`
	Nama      string    `gorm:"column:nama" json:"nama"`
	Email     string    `gorm:"column:email" json:"email"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	Password  string    `gorm:"column:password" json:"password"`
	IdRole    int       `gorm:"column:id_role" json:"id_role"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type Role struct {
	IdRole    int       `gorm:"primaryKey;column:id_role" json:"id_role"`
	NamaRole  string    `gorm:"column:nama_role" json:"nama_role"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

type JWTClaims struct {
	jwt.StandardClaims
	IdUser uint `json:"id_user"`
	IdRole int  `json:"id_role"`
}
