import React, { useState, useEffect, useCallback } from 'react';
import { useParams, Link, useNavigate } from 'react-router-dom'; // Import useNavigate
import { useAuth } from '../contexts/AuthContext';
import { getCommunityDetails, joinCommunity, leaveCommunity, deleteCommunity } from '../services/api'; // Import delete
import '../components/Community/Community.css';

const CommunityPage = () => {
  const [community, setCommunity] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [isMember, setIsMember] = useState(false);
  const [memberCount, setMemberCount] = useState(0);

  const { communityName } = useParams();
  const { user, token, isAuthenticated } = useAuth();
  const navigate = useNavigate(); // Hook untuk navigasi

  const fetchCommunityData = useCallback(async () => {
    try {
      setLoading(true);
      const data = await getCommunityDetails(communityName);
      setCommunity(data);
      setMemberCount(data.memberCount);
      // Cek apakah user yang login adalah anggota
      if (user) {
        setIsMember(data.members.some(member => member.username === user.username));
      }
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  }, [communityName, user]);

  useEffect(() => {
    fetchCommunityData();
  }, [fetchCommunityData]);

  const handleJoinToggle = async () => {
    if (!isAuthenticated) return alert('Silakan login untuk bergabung.');

    try {
      if (isMember) {
        await leaveCommunity(community.name, token);
        setIsMember(false);
        setMemberCount(prev => prev - 1);
      } else {
        await joinCommunity(community.name, token);
        setIsMember(true);
        setMemberCount(prev => prev + 1);
      }
    } catch (err) {
      console.error(err);
    }
  };

  const isCreator = user && user.userId === community?.creatorId;

    const handleDelete = async () => {
      if (!isCreator) return;
      
      // Tampilkan konfirmasi sebelum menghapus
      if (window.confirm(`Apakah Anda yakin ingin menghapus komunitas "${community.name}"? Aksi ini tidak bisa dibatalkan.`)) {
        try {
          await deleteCommunity(community.name, token);
          alert('Komunitas berhasil dihapus.');
          navigate('/'); // Arahkan ke homepage setelah berhasil
        } catch (err) {
          alert(`Gagal menghapus komunitas: ${err.message}`);
        }
      }
    };

  if (loading) return <div className="app-content"><p>Loading community...</p></div>;
  if (error) return <div className="app-content"><p style={{ color: 'red' }}>Error: {error}</p></div>;
  if (!community) return null;

  return (
    <div className="community-page">
      <div className="card community-header">
        <img src={`https://api.dicebear.com/8.x/initials/svg?seed=${community.name}`} alt="Avatar" className="community-avatar" />
        <div className="community-info">
          <h2>{community.name}</h2>
          <p>{community.description}</p>
          <span>{memberCount} anggota</span>
        </div>
        <div className="community-actions">
          {isCreator && (
            <button onClick={handleDelete} className="join-button danger">
              Hapus
            </button>
          )}
          <button onClick={handleJoinToggle} className={`join-button ${isMember ? 'secondary' : 'primary'}`}>
            {isMember ? 'Tinggalkan' : 'Gabung'}
          </button>
        </div>
      </div>

      <div className="community-body">
        <div className="community-feed">
          <h3 className="content-title">Postingan Komunitas</h3>
          <div className="card"><p>Fitur postingan khusus komunitas akan segera hadir!</p></div>
        </div>
        <div className="community-members-list">
           <h3 className="content-title">Anggota</h3>
           <div className="card">
             {community.members.map(member => (
               <Link to={`/${member.username}`} key={member.username} className="member-item">
                 <img src={member.profilePictureUrl || `https://api.dicebear.com/8.x/initials/svg?seed=${member.username}`} alt="Avatar" />
                 <span>{member.fullName}</span>
               </Link>
             ))}
           </div>
        </div>
      </div>
    </div>
  );
};

export default CommunityPage;