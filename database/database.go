package database

import (
	"BookHaven/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var db *sql.DB

func InitDB(logger *logrus.Logger) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DB_USER, config.DB_PASS, config.DB_HOST, config.DB_PORT, config.DB_NAME)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("FAILED LOAD DATABASE ", err)
	}
	err = db.Ping()
	if err != nil {
		logger.Fatal("FAILED PING DATABASE ", err)
	}
	logger.Info("SUCCESS LOAD DATABASE")
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	db.Close()
}
