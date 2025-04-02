package main

import (
	"log"
	"net/http"

	"backend/internal/database"
	"backend/internal/handlers"
)

func main() {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=postgres_l3 sslmode=disable"
	db, err := database.ConnectDB(connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")

	http.HandleFunc("/register", handlers.RegisterHandler(db))
	http.HandleFunc("/login", handlers.LoginHandler(db))

	log.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
