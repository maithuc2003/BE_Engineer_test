package database

import (
	"fmt"
	"log"
	"tucode_v2/config"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db := config.LoadDB()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh", db.Host, db.Port, db.User, db.Password, db.DBName)
	// Connect PostgreSQL database
	pgDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL database:", err)
	}
	fmt.Println("Kết nối thành công tới PostgreSQL")
	return pgDB
}
