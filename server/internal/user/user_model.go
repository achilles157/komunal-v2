package user

import "time"

// User merepresentasikan model pengguna di database
type User struct {
	ID           int64     `json:"id"`
	FullName     string    `json:"fullName"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` // Tanda "-" berarti field ini diabaikan saat encoding JSON
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
