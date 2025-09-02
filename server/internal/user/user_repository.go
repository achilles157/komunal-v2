package user

import (
	"database/sql"
	"time"
)

// UserRepository bertanggung jawab atas interaksi data User dengan database
type UserRepository struct {
	db *sql.DB
}

// FindByEmail mencari user berdasarkan alamat email
func (r *UserRepository) FindByEmail(email string) (*User, error) {
	query := `
		SELECT id, full_name, username, email, password_hash, created_at, updated_at 
		FROM users WHERE email = $1`

	var user User
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.FullName,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err // Akan mengembalikan sql.ErrNoRows jika tidak ditemukan
	}

	return &user, nil
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
