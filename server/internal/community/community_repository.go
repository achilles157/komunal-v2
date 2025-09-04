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
	// Memulai transaksi
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	// 1. Masukkan data ke tabel 'communities'
	communityQuery := `INSERT INTO communities (name, description, creator_id, created_at)
                       VALUES ($1, $2, $3, $4)
                       RETURNING id, created_at`

	now := time.Now()
	err = tx.QueryRow(communityQuery, community.Name, community.Description, community.CreatorID, now).Scan(&community.ID, &community.CreatedAt)
	if err != nil {
		tx.Rollback() // Batalkan transaksi jika gagal
		return err
	}

	// 2. Masukkan kreator sebagai admin di 'community_members'
	memberQuery := `INSERT INTO community_members (user_id, community_id, role, joined_at)
                     VALUES ($1, $2, $3, $4)`

	_, err = tx.Exec(memberQuery, community.CreatorID, community.ID, "admin", now)
	if err != nil {
		tx.Rollback() // Batalkan transaksi jika gagal
		return err
	}

	// Jika semua berhasil, commit transaksi
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
