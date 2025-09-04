package post

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type CreatePostPayload struct {
	Content  string `json:"content"`
	MediaURL string `json:"mediaUrl"`
}

type PostHandler struct {
	service *PostService
}

func NewPostHandler(service *PostService) *PostHandler {
	return &PostHandler{service: service}
}

// CreatePostHandler menangani pembuatan postingan baru
func (h *PostHandler) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	// Ambil userID dari context yang sudah diisi oleh middleware JWT
	userID, ok := r.Context().Value("userID").(int64)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
		return
	}

	var payload CreatePostPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	post, err := h.service.CreatePost(userID, payload.Content, payload.MediaURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

// GetPostsHandler menangani pengambilan semua postingan untuk feed
func (h *PostHandler) GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := h.service.GetFeedPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// GetPostsByUsernameHandler menangani pengambilan postingan milik seorang user
func (h *PostHandler) GetPostsByUsernameHandler(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	posts, err := h.service.GetPostsByUsername(username)
	if err != nil {
		http.Error(w, "Could not retrieve posts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func (h *PostHandler) LikePostHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int64)

	postIDStr := r.PathValue("postId")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	if err := h.service.LikePost(userID, postID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// PERBAIKAN: Kirim response JSON, bukan hanya status
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Post liked successfully"})
}

// UnlikePostHandler menangani permintaan untuk batal menyukai sebuah postingan
func (h *PostHandler) UnlikePostHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int64)

	postIDStr := r.PathValue("postId")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	if err := h.service.UnlikePost(userID, postID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// PERBAIKAN: Kirim response JSON, bukan hanya status
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Post unliked successfully"})
}
