package config

import (
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	// Load env
	LoadEnv()

	// Load string connection from env
	dbConf := os.Getenv("SQLSTRING")

	// Create connection
	DB, err :=  gorm.Open(mysql.Open(dbConf), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	// Set configuration database connections
	db0, err := DB.DB()
	if err != nil {
		panic(err)
	}
	db0.SetConnMaxIdleTime(time.Duration(1) * time.Minute)
	db0.SetConnMaxLifetime(time.Duration(1) * time.Minute)
	db0.SetMaxIdleConns(2)

	// Show log
	DB.Statement.RaiseErrorOnNotFound = true // Raise error on not found

	return DB
}