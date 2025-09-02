package server

import (
	"net/http"
	"time"

	"komunal/server/internal/user" // Ganti 'komunal' dengan nama modul go Anda
)

// Server adalah struct untuk server HTTP kita
type Server struct {
	server *http.Server
}

// NewServer membuat dan mengkonfigurasi server baru
func NewServer(port string, userHandler *user.UserHandler) *Server {

	// Daftarkan semua rute/endpoint di sini
	mux := http.NewServeMux()
	mux.HandleFunc("/api/register", userHandler.Register)
	// mux.HandleFunc("/api/login", userHandler.Login) // Nanti ditambahkan

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
