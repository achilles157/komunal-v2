package auth

import (
	"errors"
	"komunal/server/internal/user" // Sesuaikan dengan nama modul Anda
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// AuthService menangani logika bisnis terkait autentikasi
type AuthService struct {
	userRepo *user.UserRepository
}

// Claims adalah struktur data yang akan kita simpan di dalam JWT
type Claims struct {
	UserID   int64  `json:"userId"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// NewAuthService membuat instance baru dari AuthService
func NewAuthService(userRepo *user.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

// Login memverifikasi kredensial pengguna dan mengembalikan token jika valid
func (s *AuthService) Login(email, password string) (string, error) {
	// 1. Cari pengguna berdasarkan email
	foundUser, err := s.userRepo.FindByEmail(email)
	if err != nil {
		// Jika user tidak ditemukan atau ada error lain
		return "", errors.New("invalid email or password")
	}

	// 2. Bandingkan password yang diberikan dengan hash di database
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(password))
	if err != nil {
		// Jika password tidak cocok
		return "", errors.New("invalid email or password")
	}

	// 3. Jika kredensial valid, buat JWT
	expirationTime := time.Now().Add(24 * time.Hour) // Token berlaku 24 jam

	claims := &Claims{
		UserID:   foundUser.ID,
		Username: foundUser.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Ambil secret key dari environment variable
	jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	if len(jwtKey) == 0 {
		return "", errors.New("JWT_SECRET_KEY not set")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
