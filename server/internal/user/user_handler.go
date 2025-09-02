package user

import (
	"encoding/json"
	"log"
	"net/http"
)

// RegisterPayload adalah data yang kita harapkan dari request pendaftaran
type RegisterPayload struct {
	FullName string `json:"fullName"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserHandler struct {
	service *UserService // Menggunakan service, bukan langsung ke DB
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service: service}
}

// GetUserProfileHandler menangani request untuk mendapatkan profil user
func (h *UserHandler) GetUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Ambil username dari URL, contoh: /api/users/johndoe
	username := r.PathValue("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	userProfile, err := h.service.GetUserProfile(username)
	if err != nil {
		// Jika user tidak ditemukan, kirim status 404
		if err.Error() == "sql: no rows in result set" {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userProfile)
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var payload RegisterPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Panggil service untuk mendaftarkan user
	newUser, err := h.service.RegisterUser(payload.FullName, payload.Username, payload.Email, payload.Password)
	if err != nil {
		// Penanganan error bisa dibuat lebih canggih di sini
		log.Printf("Error registering user: %v", err)
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	// Kirim response sukses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser) // Kirim data user baru (tanpa password)
}
