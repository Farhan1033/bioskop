package infra

import (
	"bioskop/entity"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host   = "localhost"
	user   = "postgres"
	pass   = "postgres"
	dbname = "bioskop"
	port   = 5432
	DB     *gorm.DB
	err    error
)

func InitDB() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname,
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
