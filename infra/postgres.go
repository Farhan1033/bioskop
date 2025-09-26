package infra

import (
	"bioskop/entity"
	"bioskop/infra/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitDB() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.GetKey("DB_HOST"),
		config.GetKey("DB_PORT"),
		config.GetKey("DB_USER"),
		config.GetKey("DB_PASS"),
		config.GetKey("DB_NAME"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	errMigrate := DB.AutoMigrate(entity.Bioskop{})
	if errMigrate != nil {
		log.Fatal("Failed to migrate database")
	}

	fmt.Println("Database connected successfully")
}
