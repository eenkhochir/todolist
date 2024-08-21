package database

import (
	"fmt"
	"log"
	"os"
	"todolist-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initdatabase() {
	var err error
	dsn := "host=localhost user=postgres password=123 dbname=todolist port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Датабазтай холбогдоход алдаа гарлаа.")
		os.Exit(2)
	}
	fmt.Println("Датабазтай холбогдлоо.")

	// automigrate nemeh
	DB.AutoMigrate(&models.Todo{}, &models.User{})
}
