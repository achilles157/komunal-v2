package server

import (
	"encoding/json"
	"net/http"
	"time"

	"komunal/server/internal/auth"
	"komunal/server/internal/middleware"
	"komunal/server/internal/user"
)

// Server adalah struct untuk server HTTP kita
type Server struct {
	server *http.Server
}

// NewServer membuat dan mengkonfigurasi server baru
func NewServer(port string, userHandler *user.UserHandler, authHandler *auth.AuthHandler) *Server {
	mux := http.NewServeMux()

	// --- Public Routes (tidak perlu login) ---
	mux.HandleFunc("/api/register", userHandler.Register)
	mux.HandleFunc("/api/login", authHandler.Login)

	// --- Protected Routes (harus login) ---
	// Buat handler untuk profil sebagai contoh
	profileHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ambil data user dari context yang sudah diisi oleh middleware
		userID := r.Context().Value("userID").(int64)
		username := r.Context().Value("username").(string)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message":  "Welcome to your profile!",
			"userId":   userID,
			"username": username,
		})
	})

	// Terapkan middleware ke profileHandler
	mux.Handle("/api/profile", middleware.JWTAuthentication(profileHandler))

	return &Server{
		server: &http.Server{
			Addr:         ":" + port,
			Handler:      mux, // Gunakan mux yang sudah diisi rute
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
	}
}

// ListenAndServe menjalankan server
func (s *Server) ListenAndServe() error {
	return s.server.ListenAndServe()
}
