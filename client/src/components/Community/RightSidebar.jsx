import React from 'react';
import { Link } from 'react-router-dom'; // Import Link
import './Community.css';

const dummyCommunities = [
  { id: 1, name: 'Seni Digital Indonesia', members: '12k' },
  { id: 2, name: 'Fotografi Jalanan', members: '8.5k' },
  { id: 3, name: 'Ilustrator & Komikus', members: '15k' },
];

const dummyUsers = [
  { id: 1, name: 'Andi Mahardika', handle: 'seniman_kode' },
  { id: 2, name: 'Citra Lestari', handle: 'gadis_kamera' },
];

const RightSidebar = () => {
  return (
    <aside className="right-sidebar-container">
      <div className="card suggestion-card">
        <h3 className="suggestion-title">Komunitas untuk Anda</h3>
        <ul className="suggestion-list">
          {dummyCommunities.map(community => (
            // Ganti <li> dengan <Link>
            <Link to={`/c/${community.name.replace(/\s+/g, '-')}`} key={community.id} className="suggestion-item as-link">
              <img src={`https://api.dicebear.com/8.x/initials/svg?seed=${community.name}`} alt="Avatar" className="suggestion-avatar" />
              <div className="suggestion-info">
                <span className="suggestion-name">{community.name}</span>
                <span className="suggestion-meta">{community.members} anggota</span>
              </div>
            </Link>
          ))}
        </ul>
      </div>

      <div className="card suggestion-card">
        <h3 className="suggestion-title">Pengguna untuk Diikuti</h3>
        <ul className="suggestion-list">
          {dummyUsers.map(user => (
            <li key={user.id} className="suggestion-item">
              <img src={`https://api.dicebear.com/8.x/initials/svg?seed=${user.name}`} alt="User Avatar" className="suggestion-avatar" />
              <div className="suggestion-info">
                <span className="suggestion-name">{user.name}</span>
                <span className="suggestion-meta">@{user.handle}</span>
              </div>
              <button className="follow-button">Ikuti</button>
            </li>
          ))}
        </ul>
      </div>
    </aside>
  );
};

export default RightSidebar;