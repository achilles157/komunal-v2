package community

import (
	"database/sql"
	"time"
)

type CommunityRepository struct {
	db *sql.DB
}

func NewCommunityRepository(db *sql.DB) *CommunityRepository {
	return &CommunityRepository{db: db}
}

// CreateWithAdminTransaction membuat komunitas baru dan menambahkan kreator sebagai admin dalam satu transaksi
func (r *CommunityRepository) CreateWithAdminTransaction(community *Community) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	communityQuery := `INSERT INTO communities (name, slug, description, creator_id, created_at)
                       VALUES ($1, $2, $3, $4, $5)
                       RETURNING id, created_at`

	now := time.Now() // Definisikan 'now' sekali di sini
	err = tx.QueryRow(communityQuery, community.Name, community.Description, community.CreatorID, now).Scan(&community.ID, &community.CreatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}

	memberQuery := `INSERT INTO community_members (user_id, community_id, role, joined_at)
                     VALUES ($1, $2, $3, $4)`

	// Gunakan variabel 'now' yang sama untuk konsistensi
	_, err = tx.Exec(memberQuery, community.CreatorID, community.ID, "admin", now)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// FindByName mencari sebuah komunitas berdasarkan namanya yang unik
func (r *CommunityRepository) FindByName(name string) (*Community, error) {
	var community Community
	query := `SELECT id, name, description, creator_id, created_at FROM communities WHERE name = $1`
	err := r.db.QueryRow(query, name).Scan(&community.ID, &community.Name, &community.Description, &community.CreatorID, &community.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &community, nil
}

// GetMembers mengambil daftar anggota dari sebuah komunitas
func (r *CommunityRepository) GetMembers(communityID int) ([]MemberInfo, error) {
	query := `
		SELECT u.username, u.full_name, u.profile_picture_url
		FROM community_members cm
		JOIN users u ON cm.user_id = u.id
		WHERE cm.community_id = $1
		ORDER BY cm.joined_at ASC`

	rows, err := r.db.Query(query, communityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []MemberInfo
	for rows.Next() {
		var member MemberInfo
		var profilePic sql.NullString
		if err := rows.Scan(&member.Username, &member.FullName, &profilePic); err != nil {
			return nil, err
		}
		if profilePic.Valid {
			member.ProfilePictureURL = profilePic.String
		}
		members = append(members, member)
	}
	return members, nil
}

// JoinCommunity menambahkan pengguna sebagai anggota baru ke dalam komunitas
func (r *CommunityRepository) JoinCommunity(userID int64, communityID int) error {
	query := `INSERT INTO community_members (user_id, community_id, role) VALUES ($1, $2, 'member') ON CONFLICT (user_id, community_id) DO NOTHING`
	_, err := r.db.Exec(query, userID, communityID)
	return err
}

// LeaveCommunity menghapus keanggotaan pengguna dari sebuah komunitas
func (r *CommunityRepository) LeaveCommunity(userID int64, communityID int) error {
	query := `DELETE FROM community_members WHERE user_id = $1 AND community_id = $2`
	_, err := r.db.Exec(query, userID, communityID)
	return err
}

// FindByUserID mengambil semua komunitas di mana seorang pengguna adalah anggota
func (r *CommunityRepository) FindByUserID(userID int64) ([]Community, error) {
	query := `
		SELECT c.id, c.name, c.slug, c.description, c.creator_id, c.created_at
		FROM communities c
		JOIN community_members cm ON c.id = cm.community_id
		WHERE cm.user_id = $1`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var communities []Community
	for rows.Next() {
		var community Community
		var slug sql.NullString // Menangani slug yang bisa NULL
		if err := rows.Scan(&community.ID, &community.Name, &slug, &community.Description, &community.CreatorID, &community.CreatedAt); err != nil {
			return nil, err
		}
		if slug.Valid {
			community.Slug = slug.String
		}
		communities = append(communities, community)
	}
	return communities, nil
}
