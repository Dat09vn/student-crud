package database

import (
	"fmt"
	"log"
	"os"
	"student-crud/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DB_DSN") // example: "user:password@tcp(host:port)/dbname?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	db.AutoMigrate(&models.Student{})
	DB = db
	fmt.Println("Database connected!")
}
