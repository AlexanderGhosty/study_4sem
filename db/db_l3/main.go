package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Подключаемся к базе данных
	var err error
	connectDB()
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v\n", err)
	}
	defer db.Close()

	// Создаём таблицы, если их нет
	if err := initDBTables(); err != nil {
		log.Fatalf("Ошибка инициализации таблиц: %v\n", err)
	}

	// Инициализируем роутер (используем gorilla/mux для удобства)
	r := mux.NewRouter()

	// Маршруты для авторов
	r.HandleFunc("/authors", getAllAuthors).Methods("GET")
	r.HandleFunc("/authors", addAuthor).Methods("POST")
	r.HandleFunc("/authors/{id}", updateAuthor).Methods("PUT")

	// Маршруты для категорий
	r.HandleFunc("/categories", getAllCategories).Methods("GET")
	r.HandleFunc("/categories", addCategory).Methods("POST")
	r.HandleFunc("/categories/{id}", updateCategory).Methods("PUT")

	// Маршруты для книг
	r.HandleFunc("/books", getAllBooks).Methods("GET")
	r.HandleFunc("/books", addBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")

	// Маршрут для статических файлов (где лежит index.html)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))

	// Запуск сервера
	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
