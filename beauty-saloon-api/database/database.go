package database

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

// CreateConnection establishes a connection to the SQL Server database
func CreateConnection() (*sql.DB, error) {
	connString := "server=YOUR_SERVER_ADDRESS;user id=YOUR_USERNAME;password=YOUR_PASSWORD;database=YOUR_DATABASE"
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to SQL Server successfully!")
	return db, nil
}
