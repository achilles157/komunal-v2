import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { getUserProfile, getPostsByUsername } from '../services/api';
import ProfileHeader from '../components/Profile/ProfileHeader';
import ProfileNav from '../components/Profile/ProfileNav'; // 1. Import komponen Nav
import PostCard from '../components/Feed/PostCard';

const ProfilePage = () => {
  const [profile, setProfile] = useState(null);
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [activeTab, setActiveTab] = useState('posts'); // 2. State untuk tab aktif
  const { username } = useParams();

  useEffect(() => {
    const fetchProfileData = async () => {
      try {
        setLoading(true);
        const profileData = await getUserProfile(username);
        const postsData = await getPostsByUsername(username);
        
        setProfile(profileData);
        setPosts(postsData || []);
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };
    fetchProfileData();
  }, [username]);

  // Fungsi untuk merender konten berdasarkan tab yang aktif
  const renderTabContent = () => {
    switch (activeTab) {
      case 'posts':
        return (
          <>
            <h3 className="content-title">Postingan</h3>
            {posts.length > 0 ? (
              posts.map(post => (
                <PostCard key={post.id} post={{
                  ...post,
                  author: { name: post.authorName, username: post.authorUsername },
                  timestamp: new Date(post.createdAt).toLocaleString('id-ID')
                }} />
              ))
            ) : (
              <div className="card">
                <p>@{username} belum memiliki postingan.</p>
              </div>
            )}
          </>
        );
      case 'about':
        return (
          <>
            <h3 className="content-title">Tentang @{username}</h3>
            <div className="card">
              <p>{profile.bio || 'Tidak ada bio yang ditulis.'}</p>
              <hr style={{border: 'none', borderTop: '1px solid #efefef', margin: '1rem 0'}} />
              <p><strong>Bergabung pada:</strong> {new Date(profile.joinedAt).toLocaleDateString('id-ID', { year: 'numeric', month: 'long', day: 'numeric' })}</p>
            </div>
          </>
        );
      case 'creations':
        return (
          <>
            <h3 className="content-title">Karya</h3>
            <div className="card">
              <p>Fitur marketplace untuk menampilkan karya akan segera hadir!</p>
            </div>
          </>
        );
      default:
        return null;
    }
  };

  if (loading) return <div className="app-content"><p>Loading profile...</p></div>;
  if (error) return <div className="app-content"><p style={{ color: 'red' }}>Error: {error}</p></div>;
  if (!profile) return <div className="app-content"><p>User not found.</p></div>;

  return (
    <div className="profile-page-container">
      <ProfileHeader userProfile={profile} />
      {/* 3. Tambahkan Navigasi Tab di sini */}
      <ProfileNav activeTab={activeTab} setActiveTab={setActiveTab} />
      
      <div className="profile-content">
        {/* 4. Render konten dinamis berdasarkan tab */}
        {renderTabContent()}
      </div>
    </div>
  );
};

export default ProfilePage;