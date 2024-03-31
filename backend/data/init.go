package data

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	var err error
	Db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName))
	if err != nil {
		panic(err)
	}
}
