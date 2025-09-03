import React, { useState } from 'react'; // Import useState
import { Link } from 'react-router-dom';
import { useAuth } from '../../../contexts/AuthContext';
import { FiMenu, FiX } from 'react-icons/fi'; // Import ikon Menu dan X
import MobileNavMenu from './MobileNavMenu'; // Import komponen baru
import './Navbar.css';

const Navbar = () => {
  const { isAuthenticated, user, logout } = useAuth();
  const [isMenuOpen, setIsMenuOpen] = useState(false); // State untuk menu mobile

  return (
    <nav className="navbar">
      <div className="navbar-container">
        <Link to="/" className="navbar-logo">
          Komunal
        </Link>

        {/* --- Tampilan Desktop --- */}
        <div className="desktop-nav">
          {isAuthenticated && user ? (
            <div className="navbar-user-section">
              <span className="welcome-message">
                Selamat datang, <strong>{user.username}</strong>
              </span>
              <button onClick={logout} className="logout-button">
                Logout
              </button>
            </div>
          ) : (
            <div className="navbar-actions">
              <Link to="/login" className="navbar-action-button">
                Masuk
              </Link>
              <Link to="/register" className="navbar-action-button primary">
                Daftar
              </Link>
            </div>
          )}
        </div>

        {/* --- Ikon Hamburger (hanya muncul di mobile) --- */}
        <div className="mobile-nav-toggle">
          <button onClick={() => setIsMenuOpen(!isMenuOpen)} className="hamburger-button">
            {isMenuOpen ? <FiX /> : <FiMenu />}
          </button>
        </div>
      </div>

      {/* --- Menu Mobile (muncul jika isMenuOpen true) --- */}
      <MobileNavMenu isOpen={isMenuOpen} closeMenu={() => setIsMenuOpen(false)} />
    </nav>
  );
};

export default Navbar;