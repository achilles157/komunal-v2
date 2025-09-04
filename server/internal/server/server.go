package server

import (
	"encoding/json"
	"net/http"
	"time"

	"komunal/server/internal/auth"
	"komunal/server/internal/community"
	"komunal/server/internal/middleware"
	"komunal/server/internal/post"
	"komunal/server/internal/user"

	"github.com/gorilla/mux" // Import gorilla/mux
)

// Server adalah struct untuk server HTTP kita
type Server struct {
	server *http.Server
}

// NewServer membuat dan mengkonfigurasi server baru menggunakan gorilla/mux
func NewServer(port string, userHandler *user.UserHandler, authHandler *auth.AuthHandler, postHandler *post.PostHandler, communityHandler *community.CommunityHandler) *Server {
	// Gunakan mux.NewRouter() sebagai ganti http.NewServeMux()
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter() // Grup semua route di bawah /api

	// --- Rute Publik ---
	apiRouter.HandleFunc("/register", userHandler.Register).Methods("POST")
	apiRouter.HandleFunc("/login", authHandler.Login).Methods("POST")
	apiRouter.HandleFunc("/posts", postHandler.GetPostsHandler).Methods("GET")
	apiRouter.HandleFunc("/communities/{name}", communityHandler.GetCommunityHandler).Methods("GET")
	apiRouter.HandleFunc("/users/{username}", userHandler.GetUserProfileHandler).Methods("GET")
	apiRouter.HandleFunc("/users/{username}/posts", postHandler.GetPostsByUsernameHandler).Methods("GET")

	// --- Rute Terproteksi ---
	// Buat subrouter baru yang akan menggunakan middleware
	protectedRouter := apiRouter.PathPrefix("").Subrouter()
	protectedRouter.Use(middleware.JWTAuthentication) // Terapkan middleware ke grup ini

	// Definisikan handler untuk profil pribadi
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
	protectedRouter.HandleFunc("/profile", profileHandler).Methods("GET")
	protectedRouter.HandleFunc("/profile", userHandler.UpdateUserProfileHandler).Methods("PUT")

	// Rute Postingan
	protectedRouter.HandleFunc("/posts/create", postHandler.CreatePostHandler).Methods("POST")
	protectedRouter.HandleFunc("/posts/{postId}/like", postHandler.LikePostHandler).Methods("POST")
	protectedRouter.HandleFunc("/posts/{postId}/like", postHandler.UnlikePostHandler).Methods("DELETE")

	// Rute Komunitas
	protectedRouter.HandleFunc("/communities", communityHandler.CreateCommunityHandler).Methods("POST")
	protectedRouter.HandleFunc("/communities/{name}/join", communityHandler.JoinCommunityHandler).Methods("POST")
	protectedRouter.HandleFunc("/communities/{name}/join", communityHandler.LeaveCommunityHandler).Methods("DELETE")

	// Rute Follow
	protectedRouter.HandleFunc("/users/{username}/follow", userHandler.FollowUserHandler).Methods("POST")
	protectedRouter.HandleFunc("/users/{username}/follow", userHandler.UnfollowUserHandler).Methods("DELETE")

	return &Server{
		server: &http.Server{
			Addr:         ":" + port,
			Handler:      router, // Gunakan gorilla/mux router
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
