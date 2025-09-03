import React from 'react';
import './Profile.css';

const ProfileNav = ({ activeTab, setActiveTab }) => {
  return (
    <nav className="profile-nav">
      <button
        className={`nav-tab ${activeTab === 'posts' ? 'active' : ''}`}
        onClick={() => setActiveTab('posts')}
      >
        Postingan
      </button>
      <button
        className={`nav-tab ${activeTab === 'about' ? 'active' : ''}`}
        onClick={() => setActiveTab('about')}
      >
        Tentang
      </button>
      <button
        className={`nav-tab ${activeTab === 'creations' ? 'active' : ''}`}
        onClick={() => setActiveTab('creations')}
      >
        Karya
      </button>
    </nav>
  );
};

export default ProfileNav;