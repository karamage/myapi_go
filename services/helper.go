package services

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser     = "docker"
	dbPassword = "docker"
	dbDatabase = "sampledb"
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser,
		dbPassword, dbDatabase)
)

func ConnectDB() (*sql.DB, error) {
	db, _ := sql.Open("mysql", dbConn)
	err := db.Ping() // Need to do this to check that the connection is valid
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}
