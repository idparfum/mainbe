package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	IdUser 	  uint 	 	`gorm:"primaryKey;column:id_user" json:"id_user"`
	Nama 	  string 	`gorm:"colum:nama" json:"nama"`
	Email 	  string 	`gorm:"colum:email" json:"email"`
	Phone 	  string 	`gorm:"colum:phone" json:"phone"`
	Password  string 	`gorm:"colum:password" json:"password"`
	IdRole 	  int 	 	`gorm:"colum:id_role" json:"id_role"`
	CreatedAt time.Time `gorm:"colum:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"colum:updated_at" json:"updated_at"`
}

type Role struct {
	IdRole 	  int 	 	`gorm:"primaryKey;colum:id_role" json:"id_role"`
	NamaRole  string 	`gorm:"colum:nama_role" json:"nama_role"`
	CreatedAt time.Time `gorm:"colum:created_at" json:"created_at"`
}

type JWTClaims struct {
	jwt.StandardClaims
	IdUser uint `json:"id_user"`
	IdRole int `json:"id_role"`
}