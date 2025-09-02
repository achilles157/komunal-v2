import React from 'react';
import { useAuth } from '../../contexts/AuthContext';
import './Profile.css'; // Kita akan tambahkan style di file yang sudah ada

const ProfileHeader = ({ userProfile }) => {
  const { user: loggedInUser } = useAuth(); // Mendapatkan info user yang sedang login

  const isOwnProfile = loggedInUser && loggedInUser.username === userProfile.username;

  return (
    <div className="card profile-header-card">
      <div className="profile-header-info">
        <img
          src={userProfile.profilePictureUrl || `https://api.dicebear.com/8.x/initials/svg?seed=${userProfile.username}`}
          alt={`${userProfile.username}'s avatar`}
          className="profile-header-avatar"
        />
        <div className="profile-header-details">
          <h2 className="profile-header-fullname">{userProfile.fullName}</h2>
          <p className="profile-header-username">@{userProfile.username}</p>
          <p className="profile-header-bio">{userProfile.bio || 'Pengguna ini belum menulis bio.'}</p>
        </div>
        <div className="profile-header-actions">
          {isOwnProfile ? (
            <button className="action-button primary">Edit Profil</button>
          ) : (
            <button className="action-button primary">Ikuti</button>
          )}
        </div>
      </div>
    </div>
  );
};

export default ProfileHeader;