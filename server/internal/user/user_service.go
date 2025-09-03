package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// UserService menangani logika bisnis terkait User
type UserService struct {
	repo *UserRepository
}

// NewUserService membuat instance baru dari UserService
func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetUserProfile mengambil data profil publik seorang pengguna
func (s *UserService) GetUserProfile(username string, currentUserID int64) (*UserProfileResponse, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	// Panggil fungsi baru untuk mendapatkan statistik
	stats, err := s.repo.GetStatsByUserID(user.ID)
	if err != nil {
		return nil, err
	}

	isFollowing, err := s.repo.IsFollowing(currentUserID, user.ID)
	if err != nil {
		return nil, err
	}
	// Petakan data ke UserProfileResponse
	response := &UserProfileResponse{
		ID:                user.ID,
		FullName:          user.FullName,
		Username:          user.Username,
		ProfilePictureURL: user.ProfilePictureURL,
		Bio:               user.Bio,
		JoinedAt:          user.CreatedAt,
		Stats:             *stats, // Tambahkan statistik ke response
		IsFollowing:       isFollowing,
	}

	return response, nil
}

// RegisterUser memvalidasi data, melakukan hash password, dan mendaftarkan user baru
func (s *UserService) RegisterUser(fullName, username, email, password string) (*User, error) {
	// 1. Validasi input (contoh sederhana)
	if fullName == "" || username == "" || email == "" || password == "" {
		return nil, errors.New("all fields are required")
	}
	if len(password) < 8 {
		return nil, errors.New("password must be at least 8 characters long")
	}

	// 2. Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 3. Membuat objek User baru
	newUser := &User{
		FullName:     fullName,
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
	}

	// 4. Memanggil repository untuk menyimpan user
	if err := s.repo.Create(newUser); err != nil {
		// Di sini bisa ditambahkan penanganan error spesifik,
		// misalnya jika username/email sudah ada (unique constraint violation)
		return nil, err
	}

	// Mengosongkan hash password sebelum mengembalikan data
	newUser.PasswordHash = ""

	return newUser, nil
}

// UpdateProfilePayload adalah data yang kita harapkan dari client untuk update
type UpdateProfilePayload struct {
	FullName          string
	ProfilePictureURL string
	Bio               string
}

// UpdateUserProfile memvalidasi dan memperbarui profil pengguna
func (s *UserService) UpdateUserProfile(userID int64, payload UpdateProfilePayload) (*User, error) {
	// 1. Ambil data user saat ini
	//    Kita bisa membuat fungsi GetByID di repository jika perlu, tapi untuk sekarang kita bisa skip

	// 2. Validasi sederhana
	if payload.FullName == "" {
		return nil, errors.New("full name cannot be empty")
	}

	// 3. Buat objek User untuk diupdate
	userToUpdate := &User{
		ID:                userID,
		FullName:          payload.FullName,
		ProfilePictureURL: payload.ProfilePictureURL,
		Bio:               payload.Bio,
	}

	// 4. Panggil repository untuk menyimpan perubahan
	if err := s.repo.Update(userToUpdate); err != nil {
		return nil, err
	}

	return userToUpdate, nil
}

func (s *UserService) FollowUser(followerID, followingID int64) error {
	if followerID == followingID {
		return errors.New("users cannot follow themselves")
	}
	// TODO: Cek apakah user yang akan di-follow ada
	return s.repo.Follow(followerID, followingID)
}

func (s *UserService) UnfollowUser(followerID, followingID int64) error {
	return s.repo.Unfollow(followerID, followingID)
}
