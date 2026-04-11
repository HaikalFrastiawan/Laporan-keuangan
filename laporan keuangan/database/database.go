package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	// DSN: username:password@tcp(host:port)/database_name
	// Assuming root with no password on localhost:3306
	dsn := "root:@tcp(127.0.0.1:3306)/laporan_keuangan?parseTime=true"
	
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// Connection pooling configuration
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	fmt.Println("Database connection pool initialized")
	return db
}
