package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"backend/internal/models"
)

// RegisterHandler отвечает за регистрацию нового пользователя.
func RegisterHandler(db *sql.DB) http.HandlerFunc {
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

		if creds.Username == "" || creds.Email == "" || creds.Password == "" {
			http.Error(w, "Username or email or password is empty", http.StatusBadRequest)
			return
		}

		// Хэшируем пароль
		hashBytes, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		passwordHash := string(hashBytes)

		query := `INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id`
		var newId int
		err = db.QueryRow(query, creds.Username, creds.Email, passwordHash).Scan(&newId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "User created successfully",
			"user_id": newId,
		})
	}
}
