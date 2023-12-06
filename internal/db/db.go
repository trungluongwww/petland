package db

import (
	"database/sql"
	"fmt"
	"github.com/trungluongwww/petland/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var (
	db *gorm.DB
)

func Connect(cfg config.Env) *gorm.DB {
	sqlDB, err := sql.Open("mysql", dsn(cfg))
	if err != nil {
		log.Panicf("Failed connection mysql %v", err)
	}

	db, err = gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(getLoggerOpt(cfg)),
	})
	if err != nil {
		log.Panicf("Failed connection mysql %v", err)
	}

	return db
}

func GetDB() *gorm.DB {
	return db
}

func dsn(cfg config.Env) string {
	return fmt.Sprintf("%s:%s@%s/%s?parseTime=true",
		cfg.MySqlUsername,
		cfg.MySqlPassword,
		cfg.MySqlProtocol,
		cfg.MYSqlDatabase,
	)
}

func getLoggerOpt(cfg config.Env) logger.LogLevel {
	if cfg.IsDevelop() {
		return logger.Info
	}

	return logger.Error
}
