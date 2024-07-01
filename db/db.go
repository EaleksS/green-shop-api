package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	if err := godotenv.Load(); err != nil {
			log.Print("No .env file found")
	}
}

func NewPostgresStorage() (*sql.DB, error) {
	host := getEnv("HOST", "http://localhost")
	port := 5432
	user := "postgres"
	password := getEnv("PASSWORD", "root")
	dbname := getEnv("DATABASE", "postgres")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=require",
	host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	
	if err != nil {
		panic(err)
	}

	return db, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}