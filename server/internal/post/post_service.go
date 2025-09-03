package post

import "errors"

type PostService struct {
	repo *PostRepository
}

func NewPostService(repo *PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(userID int64, content string, mediaURL string) (*Post, error) {
	if content == "" && mediaURL == "" {
		return nil, errors.New("post content or media cannot be empty")
	}

	post := &Post{
		UserID:   userID,
		Content:  content,
		MediaURL: mediaURL,
	}

	if err := s.repo.Create(post); err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostService) GetFeedPosts() ([]PostResponse, error) {
	return s.repo.GetAll()
}

func (s *PostService) GetPostsByUsername(username string) ([]PostResponse, error) {
	return s.repo.FindByUsername(username)
}

// LikePost menangani logika untuk menyukai sebuah postingan
func (s *PostService) LikePost(userID int64, postID int64) error {
	// Di sini Anda bisa menambahkan validasi, misalnya:
	// - Cek apakah postingan dengan ID tersebut ada
	// - Cek apakah pengguna sudah pernah menyukai postingan ini sebelumnya
	return s.repo.LikePost(userID, postID)
}

// UnlikePost menangani logika untuk batal menyukai sebuah postingan
func (s *PostService) UnlikePost(userID int64, postID int64) error {
	return s.repo.UnlikePost(userID, postID)
}
