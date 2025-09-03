package user

import (
	"database/sql"
	"time"
)

// UserRepository bertanggung jawab atas interaksi data User dengan database
type UserRepository struct {
	db *sql.DB
}

// FindByUsername mencari user berdasarkan username unik mereka
func (r *UserRepository) FindByUsername(username string) (*User, error) {
	query := `
		SELECT id, full_name, username, email, profile_picture_url, bio, created_at, updated_at 
		FROM users WHERE username = $1`

	var user User
	// Gunakan sql.NullString untuk field yang bisa NULL di database
	var profilePic sql.NullString
	var bio sql.NullString

	err := r.db.QueryRow(query, username).Scan(
		&user.ID,
		&user.FullName,
		&user.Username,
		&user.Email,
		&profilePic,
		&bio,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err // Akan mengembalikan sql.ErrNoRows jika tidak ditemukan
	}

	// Set nilai dari NullString ke struct User jika valid
	if profilePic.Valid {
		user.ProfilePictureURL = profilePic.String
	}
	if bio.Valid {
		user.Bio = bio.String
	}

	return &user, nil
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

// GetStatsByUserID mengambil statistik (post, follower, following) untuk seorang user
func (r *UserRepository) GetStatsByUserID(userID int64) (*UserStats, error) {
	var stats UserStats

	// Query untuk menghitung jumlah postingan
	err := r.db.QueryRow(`SELECT COUNT(*) FROM posts WHERE user_id = $1`, userID).Scan(&stats.PostCount)
	if err != nil {
		return nil, err
	}

	// Query untuk menghitung jumlah pengikut (followers)
	err = r.db.QueryRow(`SELECT COUNT(*) FROM followers WHERE following_id = $1`, userID).Scan(&stats.FollowerCount)
	if err != nil {
		return nil, err
	}

	// Query untuk menghitung jumlah yang diikuti (following)
	err = r.db.QueryRow(`SELECT COUNT(*) FROM followers WHERE follower_id = $1`, userID).Scan(&stats.FollowingCount)
	if err != nil {
		return nil, err
	}

	return &stats, nil
}

// Update memperbarui data pengguna di database
func (r *UserRepository) Update(user *User) error {
	query := `
		UPDATE users
		SET full_name = $1, profile_picture_url = $2, bio = $3, updated_at = NOW()
		WHERE id = $4`

	_, err := r.db.Exec(query, user.FullName, user.ProfilePictureURL, user.Bio, user.ID)
	return err
}

// Follow menambahkan relasi 'mengikuti' ke database
func (r *UserRepository) Follow(followerID, followingID int64) error {
	query := `INSERT INTO followers (follower_id, following_id) VALUES ($1, $2)`
	_, err := r.db.Exec(query, followerID, followingID)
	return err
}

// Unfollow menghapus relasi 'mengikuti' dari database
func (r *UserRepository) Unfollow(followerID, followingID int64) error {
	query := `DELETE FROM followers WHERE follower_id = $1 AND following_id = $2`
	_, err := r.db.Exec(query, followerID, followingID)
	return err
}

// IsFollowing memeriksa apakah seorang pengguna sudah mengikuti pengguna lain
func (r *UserRepository) IsFollowing(followerID, followingID int64) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM followers WHERE follower_id = $1 AND following_id = $2)`
	err := r.db.QueryRow(query, followerID, followingID).Scan(&exists)
	return exists, err
}
