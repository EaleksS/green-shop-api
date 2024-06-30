package main

import (
	"database/sql"
	"log"

	"github.com/EaleksS/green-shop-api/cmd/api"
	"github.com/EaleksS/green-shop-api/db"
)

func main() {
	db, err := db.NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewApiServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}