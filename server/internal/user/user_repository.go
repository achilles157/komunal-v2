package user

import (
	"database/sql"
	"time"
)

// UserRepository bertanggung jawab atas interaksi data User dengan database
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository membuat instance baru dari UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create menyimpan user baru ke dalam database
func (r *UserRepository) Create(user *User) error {
	query := `
		INSERT INTO users (full_name, username, email, password_hash, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`
	
	now := time.Now()
	
	err := r.db.QueryRow(
		query,
		user.FullName,
		user.Username,
		user.Email,
		user.PasswordHash, // Password yang sudah di-hash
		now,
		now,
	).Scan(&user.ID) // Mengambil ID yang baru dibuat dan memasukkannya ke struct user

	if err != nil {
		return err
	}

	return nil
}