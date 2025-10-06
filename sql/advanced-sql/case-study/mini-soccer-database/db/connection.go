package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=!@#$%^ host=localhost port=5432 dbname=mini_soccer_sl sslmode=disable")
	if err != nil {
		log.Fatal("Error: ", err)
	}

	return db
}
