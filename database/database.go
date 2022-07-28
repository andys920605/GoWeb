package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	maxOpenConns    = 30
	connMaxLifetime = 120
	maxIdleConns    = 10
	connMaxIdleTime = 20
)

func NewDb() (*gorm.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s search_path=%s",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_DATABASE"),
		os.Getenv("PG_SSLMODE"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_SCHEAM"),
	)
	db, err := gorm.Open(os.Getenv("PG_DRIVER"), dataSourceName)
	if err != nil {
		fmt.Println("error conn postgres")
		return nil, err
	}
	db.LogMode(true)
	// SetMaxOpenConns 設定打開資料庫連接最大數量
	db.DB().SetMaxOpenConns(maxOpenConns)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused. 設定連接可重覆使用的最大時間
	// Expired connections may be closed lazily before reuse.
	db.DB().SetConnMaxLifetime(connMaxLifetime * time.Second)
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	// If MaxOpenConns is greater than 0 but less than the new MaxIdleConns, then the new MaxIdleConns will be reduced to match the MaxOpenConns limit.
	db.DB().SetMaxIdleConns(maxIdleConns)
	db.DB().SetConnMaxIdleTime(connMaxIdleTime * time.Second)
	log.Println("DB:OK")
	return db, nil
}
