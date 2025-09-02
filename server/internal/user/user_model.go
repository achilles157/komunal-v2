package user

import "time"

// User merepresentasikan model pengguna di database
type User struct {
	ID                int64     `json:"id"`
	FullName          string    `json:"fullName"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	PasswordHash      string    `json:"-"`
	ProfilePictureURL string    `json:"profilePictureUrl,omitempty"`
	Bio               string    `json:"bio,omitempty"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

// UserProfileResponse adalah data yang aman untuk ditampilkan ke publik
type UserProfileResponse struct {
	ID                int64     `json:"id"`
	FullName          string    `json:"fullName"`
	Username          string    `json:"username"`
	ProfilePictureURL string    `json:"profilePictureUrl,omitempty"`
	Bio               string    `json:"bio,omitempty"`
	JoinedAt          time.Time `json:"joinedAt"`
}
