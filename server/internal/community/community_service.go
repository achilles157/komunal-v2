package community

import "errors"

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
		Description: description,
		CreatorID:   creatorID,
	}

	if err := s.repo.CreateWithAdminTransaction(community); err != nil {
		return nil, err
	}
	return community, nil
}
