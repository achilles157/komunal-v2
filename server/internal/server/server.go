package server

import (
	"encoding/json"
	"net/http"
	"time"

	"komunal/server/internal/auth"
	"komunal/server/internal/middleware"
	"komunal/server/internal/post"
	"komunal/server/internal/user"
)

// Server adalah struct untuk server HTTP kita
type Server struct {
	server *http.Server
}

// NewServer membuat dan mengkonfigurasi server baru
func NewServer(port string, userHandler *user.UserHandler, authHandler *auth.AuthHandler, postHandler *post.PostHandler) *Server {
	mux := http.NewServeMux()

	// --- Rute Publik ---
	mux.HandleFunc("/api/register", userHandler.Register)
	mux.HandleFunc("/api/login", authHandler.Login)
	mux.HandleFunc("/api/posts", postHandler.GetPostsHandler)

	// ROUTE UNTUK PROFIL PENGGUNA
	mux.HandleFunc("GET /api/users/{username}", userHandler.GetUserProfileHandler)

	// ROUTE UNTUK POSTINGAN PENGGUNA
	mux.HandleFunc("GET /api/users/{username}/posts", postHandler.GetPostsByUsernameHandler)

	// --- Rute Terproteksi ---

	// Definisikan handler untuk profil
	profileHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

	// Terapkan middleware ke handler pembuatan post
	mux.Handle("/api/posts/create", middleware.JWTAuthentication(http.HandlerFunc(postHandler.CreatePostHandler)))

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
