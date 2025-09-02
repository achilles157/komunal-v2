package post

import (
	"database/sql"
	"time"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

// Create menyisipkan postingan baru ke dalam database
func (r *PostRepository) Create(post *Post) error {
	query := `INSERT INTO posts (user_id, content, media_url, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5)
              RETURNING id, created_at, updated_at`

	now := time.Now()
	// mediaURL bisa string kosong jika tidak ada media
	err := r.db.QueryRow(query, post.UserID, post.Content, post.MediaURL, now, now).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
	return err
}

// GetAll mengambil semua postingan dari database dengan informasi pembuatnya
func (r *PostRepository) GetAll() ([]PostResponse, error) {
	query := `
		SELECT p.id, p.content, p.media_url, p.created_at, u.full_name, u.username
		FROM posts p
		JOIN users u ON p.user_id = u.id
		ORDER BY p.created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []PostResponse
	for rows.Next() {
		var post PostResponse
		var mediaURL sql.NullString // Menangani media_url yang bisa NULL

		if err := rows.Scan(&post.ID, &post.Content, &mediaURL, &post.CreatedAt, &post.AuthorName, &post.AuthorUsername); err != nil {
			return nil, err
		}

		if mediaURL.Valid {
			post.MediaURL = mediaURL.String
		}

		posts = append(posts, post)
	}

	return posts, nil
}

// FindByUsername mengambil semua postingan yang dibuat oleh user tertentu
func (r *PostRepository) FindByUsername(username string) ([]PostResponse, error) {
	query := `
		SELECT p.id, p.content, p.media_url, p.created_at, u.full_name, u.username
		FROM posts p
		JOIN users u ON p.user_id = u.id
		WHERE u.username = $1
		ORDER BY p.created_at DESC`

	rows, err := r.db.Query(query, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []PostResponse
	for rows.Next() {
		var post PostResponse
		var mediaURL sql.NullString

		if err := rows.Scan(&post.ID, &post.Content, &mediaURL, &post.CreatedAt, &post.AuthorName, &post.AuthorUsername); err != nil {
			return nil, err
		}

		if mediaURL.Valid {
			post.MediaURL = mediaURL.String
		}

		posts = append(posts, post)
	}

	return posts, nil
}
