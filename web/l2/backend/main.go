package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// Обработчик запроса /products

func productHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	product := []Product{
		{ID: 1, Title: "Ноутбук", Description: "Мощный ноутбук с 16 ГБ ОЗУ", Price: 999.99},
		{ID: 2, Title: "Смартфон", Description: "Смартфон с большим дисплеем", Price: 499.00},
		{ID: 3, Title: "Планшет", Description: "Компактный планшет для чтения и серфинга", Price: 299.99},
		{ID: 4, Title: "Монитор", Description: "24-дюймовый Full HD монитор", Price: 199.50},
		{ID: 5, Title: "Клавиатура", Description: "Механическая клавиатура с подсветкой", Price: 89.90},
		{ID: 6, Title: "Мышь", Description: "Игровая мышь с настраиваемыми кнопками", Price: 59.99},
		{ID: 7, Title: "Наушники", Description: "Беспроводные наушники с шумоподавлением", Price: 129.99},
	}
	

	jsonData, err := json.Marshal(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func main() {
	http.HandleFunc("/products", productHandler)

	log.Println("Starting server... at http://localhost:8080/products")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
