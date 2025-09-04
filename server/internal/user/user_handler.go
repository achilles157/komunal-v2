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
	username := r.PathValue("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	// Coba ambil ID pengguna yang sedang login dari context.
	// Jika tidak ada (pengguna tamu), userID akan menjadi 0.
	var currentUserID int64 = 0
	if userID, ok := r.Context().Value("userID").(int64); ok {
		currentUserID = userID
	}

	// Panggil service dengan dua argumen
	userProfile, err := h.service.GetUserProfile(username, currentUserID)
	if err != nil {
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

// UpdateUserProfileHandler menangani permintaan untuk memperbarui profil
func (h *UserHandler) UpdateUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Ambil userID dari context, ini memastikan hanya user yang login yang bisa edit
	userID, ok := r.Context().Value("userID").(int64)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var payload UpdateProfilePayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedUser, err := h.service.UpdateUserProfile(userID, payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

// FollowUserHandler menangani permintaan untuk mengikuti user
func (h *UserHandler) FollowUserHandler(w http.ResponseWriter, r *http.Request) {
	currentUserID := r.Context().Value("userID").(int64)
	targetUsername := r.PathValue("username")

	// Dapatkan ID user yang akan di-follow
	targetUser, err := h.service.GetUserProfile(targetUsername, 0) // currentUserID bisa 0 karena tidak relevan di sini
	if err != nil {
		http.Error(w, "User to follow not found", http.StatusNotFound)
		return
	}

	if err := h.service.FollowUser(currentUserID, targetUser.ID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// PERBAIKAN: Kirim response JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User followed successfully"})
}

func (h *UserHandler) UnfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	currentUserID := r.Context().Value("userID").(int64)
	targetUsername := r.PathValue("username")

	targetUser, err := h.service.GetUserProfile(targetUsername, 0)
	if err != nil {
		http.Error(w, "User to unfollow not found", http.StatusNotFound)
		return
	}

	if err := h.service.UnfollowUser(currentUserID, targetUser.ID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// PERBAIKAN: Kirim response JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User unfollowed successfully"})
}
