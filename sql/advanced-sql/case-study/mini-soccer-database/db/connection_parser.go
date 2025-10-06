package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Connection struct {
	Host     string
	Username string
	Dbname   string
	Password string
	Port     string
}

func ConnectionParser() *Connection {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error when load env: ", err)
	}

	host := os.Getenv("DATABASE_HOST")

	username := os.Getenv("DATABASE_USERNAME")

	dbname := os.Getenv("DATABASE_DBNAME")

	password := os.Getenv("DATABASE_PASSWORD")

	port := os.Getenv("DATABASE_PORT")

	return &Connection{
		Host:     host,
		Username: username,
		Password: password,
		Dbname:   dbname,
		Port:     port,
	}

}
