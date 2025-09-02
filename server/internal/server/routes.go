package server

import (
	"net/http"
	"server/internal/user" // Sesuaikan dengan nama modul go Anda
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	userHandler := user.NewUserHandler()

	mux.HandleFunc("/api/register", userHandler.Register)
	// TODO: Tambahkan endpoint untuk login di sini "/api/login"

	return mux
}