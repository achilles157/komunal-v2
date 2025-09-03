import React from 'react';
import { Link } from 'react-router-dom'; 
import { useAuth } from '../../contexts/AuthContext';
import './Profile.css';

const ProfileHeader = ({ userProfile }) => {
  const { user: loggedInUser } = useAuth();
  const isOwnProfile = loggedInUser && loggedInUser.username === userProfile.username;

  return (
    <div className="card profile-header-card">
      <div className="profile-header-info">
        <img
          src={userProfile.profilePictureUrl || `https://api.dicebear.com/8.x/initials/svg?seed=${userProfile.username}`}
          alt={`${userProfile.username}'s avatar`}
          className="profile-header-avatar"
        />
        {/* SEMUA DETAIL, TERMASUK TOMBOL, SEKARANG ADA DI DALAM SATU WADAH INI */}
        <div className="profile-header-details">
          <h2 className="profile-header-fullname">{userProfile.fullName}</h2>
          <p className="profile-header-username">@{userProfile.username}</p>
          
          <div className="profile-stats">
            <div className="stat-item">
              <strong>{userProfile.stats.postCount}</strong>
              <span>Postingan</span>
            </div>
            <div className="stat-item">
              <strong>{userProfile.stats.followerCount}</strong>
              <span>Pengikut</span>
            </div>
            <div className="stat-item">
              <strong>{userProfile.stats.followingCount}</strong>
              <span>Mengikuti</span>
            </div>
          </div>
          
          <p className="profile-header-bio">{userProfile.bio || 'Pengguna ini belum menulis bio.'}</p>

          {/* === BAGIAN YANG DIPINDAHKAN === */}
          <div className="profile-header-actions">
            {isOwnProfile ? (
                <Link to="/settings/profile" className="action-button primary">
                  Edit Profil
                </Link>
              ) : (
                <button className="action-button primary">Ikuti</button>
              )}
          </div>
        </div>
      </div>
    </div>
  );
};

export default ProfileHeader;