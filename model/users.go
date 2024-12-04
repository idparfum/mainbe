package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	IdUser uint `gorm:"primaryKey;colum:id_user" json:"id_user"`
	Nama string `gorm:"colum:nama" json:"nama"`
	Email string `gorm:"colum:email" json:"email"`
	Phone string `gorm:"colum:phone" json:"phone"`
	Password string `gorm:"colum:password" json:"password"`
	IdRole uint `gorm:"colum:id_role" json:"id_role"`
	CreatedAt string `gorm:"colum:created_at" json:"created_at"`
	UpdatedAt string `gorm:"colum:updated_at" json:"updated_at"`
}

type Role struct {
	IdRole uint `gorm:"primaryKey;colum:id_role" json:"id_role"`
	NamaRole string `gorm:"colum:nama_role" json:"nama_role"`
	CreatedAt string `gorm:"colum:created_at" json:"created_at"`
}

type JWTClaims struct {
	jwt.StandardClaims
	IdUser uint `json:"id_user"`
	IdRole uint `json:"id_role"`
}