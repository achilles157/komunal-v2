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
