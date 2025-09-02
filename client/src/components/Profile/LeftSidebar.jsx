import React from 'react';
import { NavLink } from 'react-router-dom';
import { useAuth } from '../../contexts/AuthContext';
import { FiHome, FiCompass, FiShoppingBag } from 'react-icons/fi';
import './Profile.css';

const LeftSidebar = () => {
  const { isAuthenticated, user } = useAuth();

  return (
    <aside className="left-sidebar-container">
      {isAuthenticated && user && (
        <div className="card profile-card">
          <img 
            src={`https://api.dicebear.com/8.x/initials/svg?seed=${user.username}`} 
            alt="User Avatar" 
            className="profile-avatar" 
          />
          <h3 className="profile-name">{user.fullName || user.username}</h3>
          <p className="profile-username">@{user.username}</p>
          <NavLink to="/profile" className="view-profile-button">
            Lihat Profil
          </NavLink>
        </div>
      )}

      <div className="card nav-card">
        <nav className="sidebar-nav">
          <NavLink to="/" className="nav-item">
            <FiHome /> Beranda
          </NavLink>
          <NavLink to="/explore" className="nav-item">
            <FiCompass /> Explore
          </NavLink>
          <NavLink to="/marketplace" className="nav-item">
            <FiShoppingBag /> Marketplace
          </NavLink>
        </nav>
      </div>
    </aside>
  );
};

export default LeftSidebar;