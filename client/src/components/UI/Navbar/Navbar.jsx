import React from 'react';
import { Link } from 'react-router-dom'; // Import Link
import { useAuth } from '../../../contexts/AuthContext';
import './Navbar.css';

const Navbar = () => {
  const { isAuthenticated, user, logout } = useAuth();

  return (
    <nav className="navbar">
      <div className="navbar-container">
        <Link to="/" className="navbar-logo">
          Komunal
        </Link>

        {isAuthenticated && user ? (
          // TAMPILAN JIKA SUDAH LOGIN DAN DATA USER SUDAH SIAP
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
    </nav>
  );
};

export default Navbar;