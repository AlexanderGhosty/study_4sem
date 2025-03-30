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
	}

	jsonData, err := json.Marshal(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func main() {
	http.HandleFunc("/product", productHandler)

	log.Println("Starting server... at http://localhost:8080/product")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
