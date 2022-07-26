package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

var DB *sql.DB

func NewDb() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")

	//組合sql連線字串
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbDatabase)
	//連接MySQL
	DB, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return
	}
	DB.SetConnMaxLifetime(time.Duration(10) * time.Second)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
}
