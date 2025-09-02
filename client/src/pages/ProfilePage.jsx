import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { getUserProfile, getPostsByUsername } from '../services/api'; // Import service baru
import ProfileHeader from '../components/Profile/ProfileHeader';
import PostCard from '../components/Feed/PostCard';

const ProfilePage = () => {
  const [profile, setProfile] = useState(null);
  const [posts, setPosts] = useState([]); // State ini akan kita gunakan sekarang
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const { username } = useParams();

  useEffect(() => {
    const fetchProfileData = async () => {
      try {
        setLoading(true);
        // Ambil data profil dan postingan secara bersamaan untuk efisiensi
        const profileData = await getUserProfile(username);
        const postsData = await getPostsByUsername(username);
        
        setProfile(profileData);
        setPosts(postsData || []); // Pastikan posts adalah array
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchProfileData();
  }, [username]);

  if (loading) return <div className="app-content"><p>Loading profile...</p></div>;
  if (error) return <div className="app-content"><p style={{ color: 'red' }}>Error: {error}</p></div>;
  if (!profile) return <div className="app-content"><p>User not found.</p></div>;

  return (
    <div className="profile-page-container">
      <ProfileHeader userProfile={profile} />
      <div className="profile-content">
        <div className="profile-posts">
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
        </div>
        <div className="profile-sidebar">
          {/* Placeholder */}
        </div>
      </div>
    </div>
  );
};

export default ProfilePage;