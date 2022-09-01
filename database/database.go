package database

import (
	model_com "GoWeb/models/commons"
	"fmt"
	"log"
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

func NewDb(Opt *model_com.Options) (*gorm.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s search_path=%s",
		Opt.Config.Postgres.Host,
		Opt.Config.Postgres.Port,
		Opt.Config.Postgres.User,
		Opt.Config.Postgres.Dbname,
		Opt.Config.Postgres.SSLMode,
		Opt.Config.Postgres.Password,
		Opt.Config.Postgres.Schema,
	)
	db, err := gorm.Open(Opt.Config.Postgres.PgDriver, dataSourceName)
	if err != nil {
		log.Printf("Conn postgres err message:%e", err)
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
	log.Println("Postgres connected")
	return db, nil
}
