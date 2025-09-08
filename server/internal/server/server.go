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

	"github.com/gorilla/handlers" // 1. Import gorilla/handlers
	"github.com/gorilla/mux"
)

// Server adalah struct untuk server HTTP kita
type Server struct {
	server *http.Server
}

// NewServer membuat dan mengkonfigurasi server baru menggunakan gorilla/mux
func NewServer(port string, userHandler *user.UserHandler, authHandler *auth.AuthHandler, postHandler *post.PostHandler, communityHandler *community.CommunityHandler) *Server {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()

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
	// PERBAIKAN: Gunakan router.Handle() untuk setiap rute terproteksi dan bungkus dengan middleware
	apiRouter.Handle("/profile", middleware.JWTAuthentication(profileHandler)).Methods("GET")
	apiRouter.Handle("/profile", middleware.JWTAuthentication(http.HandlerFunc(userHandler.UpdateUserProfileHandler))).Methods("PUT")

	// Rute Postingan Terproteksi
	apiRouter.Handle("/posts/create", middleware.JWTAuthentication(http.HandlerFunc(postHandler.CreatePostHandler))).Methods("POST")
	apiRouter.Handle("/posts/{postId}/like", middleware.JWTAuthentication(http.HandlerFunc(postHandler.LikePostHandler))).Methods("POST")
	apiRouter.Handle("/posts/{postId}/like", middleware.JWTAuthentication(http.HandlerFunc(postHandler.UnlikePostHandler))).Methods("DELETE")

	// Rute Komunitas Terproteksi
	apiRouter.Handle("/communities", middleware.JWTAuthentication(http.HandlerFunc(communityHandler.CreateCommunityHandler))).Methods("POST")
	apiRouter.Handle("/communities/{name}/join", middleware.JWTAuthentication(http.HandlerFunc(communityHandler.JoinCommunityHandler))).Methods("POST")
	apiRouter.Handle("/communities/{name}/join", middleware.JWTAuthentication(http.HandlerFunc(communityHandler.LeaveCommunityHandler))).Methods("DELETE")
	apiRouter.Handle("/user/communities", middleware.JWTAuthentication(http.HandlerFunc(communityHandler.GetUserCommunitiesHandler))).Methods("GET")

	// Rute Follow Terproteksi
	apiRouter.Handle("/users/{username}/follow", middleware.JWTAuthentication(http.HandlerFunc(userHandler.FollowUserHandler))).Methods("POST")
	apiRouter.Handle("/users/{username}/follow", middleware.JWTAuthentication(http.HandlerFunc(userHandler.UnfollowUserHandler))).Methods("DELETE")

	// --- Konfigurasi CORS ---
	// 2. Definisikan dari mana saja request diizinkan
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:5173"}) // Ganti port jika frontend Anda berbeda
	// 3. Definisikan header apa saja yang diizinkan
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	// 4. Definisikan metode HTTP apa saja yang diizinkan
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	return &Server{
		server: &http.Server{
			Addr: ":" + port,
			// 5. Bungkus router utama dengan CORS handler
			Handler:      handlers.CORS(allowedOrigins, allowedHeaders, allowedMethods)(router),
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
