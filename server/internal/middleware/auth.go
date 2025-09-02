package middleware

import (
	"context"
	"komunal/server/internal/auth" // Sesuaikan dengan nama modul Anda
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// JWTAuthentication adalah middleware untuk memverifikasi token JWT
func JWTAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Dapatkan token dari header Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Header biasanya formatnya "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}
		tokenString := tokenParts[1]

		// 2. Parse dan validasi token
		claims := &auth.Claims{}
		jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// 3. Jika valid, simpan informasi user di context request
		// Ini memungkinkan handler selanjutnya untuk mengetahui siapa user yang melakukan request
		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
		ctx = context.WithValue(ctx, "username", claims.Username)

		// Lanjutkan ke handler berikutnya dengan context yang sudah diperbarui
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
