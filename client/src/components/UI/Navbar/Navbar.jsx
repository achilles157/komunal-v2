import React from 'react';
import { useAuth } from '../../../contexts/AuthContext'; // Import useAuth
import './Navbar.css';

// Anda bisa mengganti ini dengan ikon search dari library seperti react-icons
const SearchIcon = () => (
  <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
    <circle cx="11" cy="11" r="8"></circle>
    <line x1="21" y1="21" x2="16.65" y2="16.65"></line>
  </svg>
);


const Navbar = () => {
  const { isAuthenticated, user, logout } = useAuth(); // Ambil status otentikasi, data user, dan fungsi logout

  return (
    <nav className="navbar">
      <div className="navbar-container">
        <a href="/" className="navbar-logo">
          Komunal
        </a>

        {isAuthenticated ? (
          // TAMPILAN JIKA SUDAH LOGIN
          <div className="navbar-user-section">
            <span className="welcome-message">
              Selamat datang, <strong>{user.username}</strong>
            </span>
            <button onClick={logout} className="logout-button">
              Logout
            </button>
          </div>
        ) : (
          // TAMPILAN JIKA BELUM LOGIN
          <>
            <div className="navbar-links">
              <a href="/" className="nav-link active">Beranda</a>
              <a href="/explore" className="nav-link">Eksplorasi</a>
              <a href="/marketplace" className="nav-link">Marketplace</a>
            </div>
            <div className="navbar-search">
              <SearchIcon />
              <input type="text" placeholder="Cari..." />
            </div>
          </>
        )}
      </div>
    </nav>
  );
};

export default Navbar;