package post

import "time"

// Post merepresentasikan tabel 'posts' di database
type Post struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"userId"`
	Content   string    `json:"content"`
	MediaURL  string    `json:"mediaUrl,omitempty"` // omitempty agar tidak muncul jika kosong
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// PostResponse adalah struct untuk mengirim data post ke client, termasuk info author
type PostResponse struct {
	ID             int64     `json:"id"`
	Content        string    `json:"content"`
	MediaURL       string    `json:"mediaUrl,omitempty"`
	CreatedAt      time.Time `json:"createdAt"`
	AuthorName     string    `json:"authorName"`
	AuthorUsername string    `json:"authorUsername"`
}
