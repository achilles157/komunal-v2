import React, { useState, useEffect } from 'react'; // Import useState & useEffect
import { NavLink, Link } from 'react-router-dom';
import { useAuth } from '../../contexts/AuthContext';
import { getUserCommunities } from '../../services/api'; // Import service baru
import { FiHome, FiCompass, FiShoppingBag } from 'react-icons/fi';
import './Profile.css';

const LeftSidebar = () => {
  const { isAuthenticated, user, token } = useAuth();
  const [myCommunities, setMyCommunities] = useState([]);

  useEffect(() => {
    // Ambil data komunitas hanya jika pengguna sudah login
    if (isAuthenticated && token) {
      getUserCommunities(token)
        .then(data => {
          setMyCommunities(data || []);
        })
        .catch(console.error);
    }
  }, [isAuthenticated, token]); // Jalankan ulang jika status login berubah

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
          <NavLink to={`/${user.username}`} className="view-profile-button">
            Lihat Profil
          </NavLink>
        </div>
      )}

      {isAuthenticated && (
         <Link to="/new/community" className="create-community-button">
            Buat Komunitas
         </Link>
      )}

      {isAuthenticated && myCommunities.length > 0 && (
        <div className="card my-communities-card">
          <h4 className="my-communities-title">Komunitas Saya</h4>
          <div className="my-communities-list">
            {myCommunities.map(community => (
              <Link to={`/c/${community.slug || community.name}`} key={community.id} className="community-item">
                <img src={`https://api.dicebear.com/8.x/initials/svg?seed=${community.name}`} alt="Avatar" />
                <span>{community.name}</span>
              </Link>
            ))}
          </div>
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