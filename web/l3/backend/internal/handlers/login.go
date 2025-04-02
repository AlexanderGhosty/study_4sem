package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"backend/internal/models"
	"backend/internal/token"
)

// LoginHandler отвечает за авторизацию пользователя.
func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		var creds models.Credentials
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, "Incorrect JSON", http.StatusBadRequest)
			return
		}

		if creds.Email == "" || creds.Password == "" {
			http.Error(w, "Email or password is empty", http.StatusBadRequest)
			return
		}

		var id int
		var dbUsername, dbEmail, dbPassword string
		err := db.QueryRow("SELECT id, username, email, password_hash FROM users WHERE email = $1", creds.Email).
			Scan(&id, &dbUsername, &dbEmail, &dbPassword)

		if err == sql.ErrNoRows {
			http.Error(w, "Incorrect email or password", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Проверяем хэш пароля
		err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(creds.Password))
		if err != nil {
			http.Error(w, "Incorrect email or password", http.StatusNotFound)
			return
		}

		// Генерируем JWT
		tokenString, err := token.GenerateToken(id, dbUsername)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Возвращаем клиенту токен
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"token": tokenString,
		})
	}
}
