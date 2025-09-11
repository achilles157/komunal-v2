package community

import "time"

// Community merepresentasikan tabel 'communities'
type Community struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"` // <-- TAMBAHKAN BARIS INI
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

// MemberInfo adalah struct untuk menampilkan info anggota di halaman komunitas
type MemberInfo struct {
	Username          string `json:"username"`
	FullName          string `json:"fullName"`
	ProfilePictureURL string `json:"profilePictureUrl,omitempty"`
}

// CommunityDetailsResponse adalah struct lengkap yang dikirim ke client
type CommunityDetailsResponse struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	CreatorID   int64        `json:"creatorId"`
	CreatedAt   time.Time    `json:"createdAt"`
	Members     []MemberInfo `json:"members"`
	MemberCount int          `json:"memberCount"`
}
