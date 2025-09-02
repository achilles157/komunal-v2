package user

import "time"

// User merepresentasikan model pengguna di database
type User struct {
	ID        int64     `json:"id"`
	FullName  string    `json:"fullName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Jangan pernah mengirim password hash ke client
	CreatedAt time.Time `json:"createdAt"`
}