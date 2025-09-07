package community

import (
	"errors"
	"regexp" // Import paket regex
	"strings"
)

type CommunityService struct {
	repo *CommunityRepository
}

func NewCommunityService(repo *CommunityRepository) *CommunityService {
	return &CommunityService{repo: repo}
}

func (s *CommunityService) CreateCommunity(name, description string, creatorID int64) (*Community, error) {
	if name == "" {
		return nil, errors.New("community name cannot be empty")
	}

	community := &Community{
		Name:        name,
		Slug:        generateSlug(name), // Buat slug dari nama
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

// Fungsi helper untuk membuat slug
func generateSlug(name string) string {
	// Ganti spasi dan karakter non-alfanumerik dengan strip
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	slug := reg.ReplaceAllString(strings.ToLower(name), "-")
	return strings.Trim(slug, "-")
}

func (s *CommunityService) GetUserCommunities(userID int64) ([]Community, error) {
	return s.repo.FindByUserID(userID)
}
