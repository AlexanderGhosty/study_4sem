package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"db_l4/internal/storage"
)

// GetAllItems обрабатывает GET-запрос (GET /items)
func GetAllItems(store storage.Storage, w http.ResponseWriter, r *http.Request) {
	items, err := store.GetAll()
	if err != nil {
		http.Error(w, "Ошибка при получении данных", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(items)
}

// CreateItem обрабатывает POST-запрос (POST /items)
func CreateItem(store storage.Storage, w http.ResponseWriter, r *http.Request) {
	var newItem storage.Item
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, "Невалидный JSON", http.StatusBadRequest)
		return
	}

	created, err := store.AddItem(newItem)
	if err != nil {
		http.Error(w, "Ошибка при добавлении элемента", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(created)
}

// UpdateItem обрабатывает PUT-запрос (PUT /items/{id})
func UpdateItem(store storage.Storage, w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "ID не указан", http.StatusBadRequest)
		return
	}
	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный формат ID", http.StatusBadRequest)
		return
	}

	var updatedData storage.Item
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		http.Error(w, "Невалидный JSON", http.StatusBadRequest)
		return
	}

	if err := store.UpdateItem(id, updatedData); err != nil {
		if err.Error() == "элемент не найден" {
			http.Error(w, "Элемент не найден", http.StatusNotFound)
			return
		}
		http.Error(w, "Ошибка при обновлении", http.StatusInternalServerError)
		return
	}

	items, err := store.GetAll()
	if err != nil {
		http.Error(w, "Ошибка при получении обновлённого списка", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(items)
}
