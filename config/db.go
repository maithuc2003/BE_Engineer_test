// package config

// import (
// 	"database/sql"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// 	_ "github.com/lib/pq"
// )

// func ConnectDB() *sql.DB {
// 	mySQLDB, err := sql.Open("mysql", "root@tcp(localhost:3306)/robot_db?parseTime=true&loc=Asia%2FHo_Chi_Minh")
// 	if err != nil {
// 		log.Fatal("Khong the ket noi duoc", err)
// 	}
// 	if err := mySQLDB.Ping(); err != nil {
// 		log.Fatal("Khong the ping toi DB", err)

//		}
//		log.Println("Ket noi thanh cong")
//		return mySQLDB
//	}
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func LoadDB() DBConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	return DBConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
	}

}
