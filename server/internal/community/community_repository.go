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
