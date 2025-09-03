import React from 'react';
import { NavLink, Link } from 'react-router-dom';
import { useAuth } from '../../../contexts/AuthContext';
import { FiHome, FiCompass, FiShoppingBag, FiPlusSquare } from 'react-icons/fi';

const MobileNavMenu = ({ isOpen, closeMenu }) => {
  const { isAuthenticated, user, logout } = useAuth();

  const handleLogout = () => {
    logout();
    closeMenu();
  };

  return (
    <div className={`mobile-nav-menu ${isOpen ? 'open' : ''}`}>
      <nav className="mobile-nav-links">
        {isAuthenticated && user && (
          // Bungkus seluruh div ini dengan Link
          <Link to={`/${user.username}`} className="mobile-profile-summary" onClick={closeMenu}>
            <img 
              src={user.profilePictureUrl || `https://api.dicebear.com/8.x/initials/svg?seed=${user.username}`} 
              alt="Avatar" 
              className="avatar" 
            />
            <div className="profile-info">
              <strong>{user.fullName || user.username}</strong>
              <span>@{user.username}</span>
            </div>
          </Link>
        )}

        <NavLink to="/" className="nav-item" onClick={closeMenu}>
          <FiHome /> Beranda
        </NavLink>
        <NavLink to="/explore" className="nav-item" onClick={closeMenu}>
          <FiCompass /> Explore
        </NavLink>
        <NavLink to="/marketplace" className="nav-item" onClick={closeMenu}>
          <FiShoppingBag /> Marketplace
        </NavLink>
        {isAuthenticated && (
          <NavLink to="/new/community" className="nav-item" onClick={closeMenu}>
            <FiPlusSquare /> Buat Komunitas
          </NavLink>
        )}
      </nav>

      <div className="mobile-nav-actions">
        {isAuthenticated ? (
          <button onClick={handleLogout} className="logout-button-mobile">
            Logout
          </button>
        ) : (
          <>
            <Link to="/login" className="navbar-action-button" onClick={closeMenu}>
              Masuk
            </Link>
            <Link to="/register" className="navbar-action-button primary" onClick={closeMenu}>
              Daftar
            </Link>
          </>
        )}
      </div>
    </div>
  );
};

export default MobileNavMenu;