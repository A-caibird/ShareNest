package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(driverName, datasourceName string) error {
	var err error
	DB, err = sql.Open(driverName, datasourceName)
	if err != nil {
		DB = nil
		return err
	}
	// config db
	DB.SetMaxOpenConns(10000)
	DB.SetMaxIdleConns(10)
	DB.SetConnMaxLifetime(time.Hour)
	if err := DB.Ping(); err != nil {
		DB = nil
		return err
	}
	fmt.Printf("%s database connect successfully\n", driverName)
	return nil
}

func GetDB() *sql.DB {
	return DB
}
