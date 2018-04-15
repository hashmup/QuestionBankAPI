package config

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // DB Driver
	"github.com/jmoiron/sqlx"
)

// DBConfig contains the info required for DB connection.
type DBConfig struct {
	DBUserName    string
	DBPassword    string
	DBDefaultHost string
	DBDefaultPort string
	DBName        string
}

// NewConnection uses the env vars to establish db connection and returns it.
func NewDBConnection(dbConfig DBConfig) (*sqlx.DB, error) {
	var db *sqlx.DB
	// Create a connection.
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbConfig.DBUserName,
		dbConfig.DBPassword,
		dbConfig.DBDefaultHost,
		dbConfig.DBDefaultPort,
		dbConfig.DBName)
	fmt.Printf("Connecting to: %s.\n", dsn)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(5 * time.Minute)
	fmt.Printf("Succeeded to connect to db.\n")
	return db, nil
}
