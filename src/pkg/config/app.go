package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB ...

var (
	DB *gorm.DB
)

// ConnectDB ...

func ConnectDB() {
	enverr := godotenv.Load("../../../.env")
	if enverr != nil {
		log.Fatalf("Error loading .env file")
	}
	dsn := "host=" + os.Getenv("POSTGRE_HOST") + " user=" + os.Getenv("POSTGRE_USER") + " password=" + os.Getenv("POSTGRE_PASSWORD") + " dbname=" + os.Getenv("POSTGRE_DB") + " port=" + os.Getenv("POSTGRE_PORT") + " sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		log.Println("Database connected")
		DB = db
	}
}

// GetDB ...

func GetDB() *gorm.DB {
	return DB
}
