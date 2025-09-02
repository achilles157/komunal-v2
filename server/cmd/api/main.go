package main

import (
	"log"
	"net/http"
	"os"

	"komunal/server/internal/database" // Ganti 'komunal' dengan nama modul go Anda
	"komunal/server/internal/server"
	"komunal/server/internal/user"

	"github.com/joho/godotenv"
)

func main() {
	// 1. Muat environment variables dari file .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// 2. Inisialisasi koneksi database
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	// 3. Inisialisasi dan hubungkan semua lapisan (dependency injection)
	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	// 4. Inisialisasi server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Port default jika tidak diset
	}

	srv := server.NewServer(port, userHandler) // Kirim handler ke server

	// 5. Jalankan server
	log.Printf("Server starting on port %s\n", port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", port, err)
	}
}
