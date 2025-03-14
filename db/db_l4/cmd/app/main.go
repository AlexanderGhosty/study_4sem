package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"db_l4/internal/handlers"
	"db_l4/internal/storage"
)

func main() {
	// Инициализация хранилища
	store, err := storage.NewJSONStorage("data.json")
	if err != nil {
		log.Fatalf("Не удалось инициализировать хранилище: %v\n", err)
	}

	mux := http.NewServeMux()

	staticDir := http.Dir(filepath.Join("web", "static"))
	mux.Handle("/", http.FileServer(staticDir))

	mux.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetAllItems(store, w, r)
		case http.MethodPost:
			handlers.CreateItem(store, w, r)
		default:
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/items/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			handlers.UpdateItem(store, w, r)
		} else {
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Сервер запущен на http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
