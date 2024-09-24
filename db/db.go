package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var error error
	DB, error = sql.Open("mysql", "root:aA1243690.@/products?parseTime=true")
	if error != nil {
		panic("Could not connect to database")
	}
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}
