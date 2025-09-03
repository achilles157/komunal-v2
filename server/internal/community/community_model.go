package community

import "time"

// Community merepresentasikan tabel 'communities'
type Community struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatorID   int64     `json:"creatorId"`
	CreatedAt   time.Time `json:"createdAt"`
}

// CommunityMember merepresentasikan tabel 'community_members'
type CommunityMember struct {
	UserID      int64
	CommunityID int
	Role        string // 'admin', 'moderator', 'member'
}
