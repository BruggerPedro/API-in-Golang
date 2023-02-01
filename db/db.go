package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // O "_" significa que vai utilizar essa biblioteca apenas durante o tempo de execução da aplicação.
)

func ConnectDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	user_dbpost := os.Getenv("USER_DBPOST")
	dbname := os.Getenv("DBNAME")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")

	conn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", user_dbpost, dbname, password, host)
	db, err := sql.Open("postgres", conn)

	if err != nil {
		panic(err.Error())
	}
	return db
}
