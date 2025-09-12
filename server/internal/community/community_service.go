package community

import (
	"errors"
	"regexp" // <-- 1. Tambahkan import ini
)

type CommunityService struct {
	repo *CommunityRepository
}

func NewCommunityService(repo *CommunityRepository) *CommunityService {
	return &CommunityService{repo: repo}
}

func (s *CommunityService) CreateCommunity(name, slug, description string, creatorID int64) (*Community, error) {
	if name == "" || slug == "" {
		return nil, errors.New("community name and slug cannot be empty")
	}

	// Validasi format slug (hanya huruf kecil, angka, dan strip)
	match, _ := regexp.MatchString("^[a-z0-9-]+$", slug)
	if !match {
		return nil, errors.New("community username can only contain lowercase letters, numbers, and hyphens")
	}

	// Cek keunikan slug
	exists, err := s.repo.CheckSlugExists(slug)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("community username is already taken")
	}

	community := &Community{
		Name:        name,
		Slug:        slug, // Gunakan slug dari input pengguna
		Description: description,
		CreatorID:   creatorID,
	}

	if err := s.repo.CreateWithAdminTransaction(community); err != nil {
		return nil, err
	}
	return community, nil
}

// GetCommunityDetails mengambil semua informasi yang diperlukan untuk halaman detail komunitas
func (s *CommunityService) GetCommunityDetails(name string) (*CommunityDetailsResponse, error) {
	// 1. Dapatkan detail dasar komunitas
	community, err := s.repo.FindByName(name)
	if err != nil {
		return nil, err
	}

	// 2. Dapatkan daftar anggota
	members, err := s.repo.GetMembers(community.ID)
	if err != nil {
		return nil, err
	}

	// 3. Gabungkan semua data menjadi satu response object
	response := &CommunityDetailsResponse{
		ID:          community.ID,
		Name:        community.Name,
		Description: community.Description,
		CreatorID:   community.CreatorID,
		CreatedAt:   community.CreatedAt,
		Members:     members,
		MemberCount: len(members),
	}

	return response, nil
}

// JoinCommunity menangani logika untuk bergabung ke komunitas
func (s *CommunityService) JoinCommunity(userID int64, communityID int) error {
	// Anda bisa menambahkan validasi di sini, misalnya cek apakah komunitas private, dll.
	return s.repo.JoinCommunity(userID, communityID)
}

// LeaveCommunity menangani logika untuk meninggalkan komunitas
func (s *CommunityService) LeaveCommunity(userID int64, communityID int) error {
	// Anda bisa menambahkan validasi, misalnya, kreator tidak bisa meninggalkan komunitasnya sendiri.
	return s.repo.LeaveCommunity(userID, communityID)
}

func (s *CommunityService) GetUserCommunities(userID int64) ([]Community, error) {
	return s.repo.FindByUserID(userID)
}

// DeleteCommunity menangani logika untuk menghapus komunitas
func (s *CommunityService) DeleteCommunity(communityID int, userID int64) error {
	// Validasi tambahan bisa ditaruh di sini jika perlu
	return s.repo.Delete(communityID, userID)
}
