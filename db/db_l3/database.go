package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func connectDB() {
	var err error
	connStr := "host=localhost port=5432 user=postgres dbname=l4 password=postgres sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected!")
}

// initDBTables создаёт таблицы (если они не существуют)
func initDBTables() error {
	// Пример создания таблицы Authors
	createAuthorsTable := `
    CREATE TABLE IF NOT EXISTS authors (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL
    );`

	// Пример создания таблицы Categories
	createCategoriesTable := `
    CREATE TABLE IF NOT EXISTS categories (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL
    );`

	// Пример создания таблицы Books
	createBooksTable := `
    CREATE TABLE IF NOT EXISTS books (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        author_id INT REFERENCES authors(id) ON DELETE CASCADE,
        category_id INT REFERENCES categories(id) ON DELETE CASCADE
    );`

	_, err := db.Exec(createAuthorsTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(createCategoriesTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(createBooksTable)
	if err != nil {
		return err
	}

	return nil
}
