import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { useAuth } from '../../contexts/AuthContext';
import { followUser, unfollowUser } from '../../services/api'; // Import service baru
import './Profile.css';

const ProfileHeader = ({ userProfile }) => {
  const { user: loggedInUser, token } = useAuth();
  const isOwnProfile = loggedInUser && loggedInUser.username === userProfile.username;

  // State untuk mengelola status follow secara interaktif
  const [isFollowing, setIsFollowing] = useState(userProfile.isFollowing);
  const [isLoadingFollow, setIsLoadingFollow] = useState(false);

  const handleFollowToggle = async () => {
    setIsLoadingFollow(true);
    try {
      if (isFollowing) {
        await unfollowUser(userProfile.username, token);
        setIsFollowing(false);
      } else {
        await followUser(userProfile.username, token);
        setIsFollowing(true);
      }
    } catch (error) {
      console.error("Failed to toggle follow:", error);
      // Kembalikan ke state semula jika error
      setIsFollowing(!isFollowing);
    } finally {
      setIsLoadingFollow(false);
    }
  };

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
                <Link to="/settings/profile" className="action-button">
                Edit Profil
                </Link>
            ) : (
                <button 
                    className={`action-button ${isFollowing ? 'secondary' : 'primary'}`}
                    onClick={handleFollowToggle}
                    disabled={isLoadingFollow}
                >
                    {isLoadingFollow ? '...' : (isFollowing ? 'Mengikuti' : 'Ikuti')}
                </button>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};

export default ProfileHeader;