package post

import (
	"encoding/json"
	"net/http"
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
