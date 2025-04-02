package main

import (
	"log"
	"net/http"

	"backend/internal/database"
	"backend/internal/handlers"
	"backend/internal/middleware"
)

func main() {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=postgres_l3 sslmode=disable"
	db, err := database.ConnectDB(connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")

	mux := http.NewServeMux()
	mux.HandleFunc("/register", handlers.RegisterHandler(db))
	mux.HandleFunc("/login", handlers.LoginHandler(db))

	handlerWithCORS := middleware.CORSMiddleware(mux)

	log.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", handlerWithCORS); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
