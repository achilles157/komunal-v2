import React, { useState, useEffect, useCallback } from 'react';
import CreatePost from '../components/Feed/CreatePost';
import PostCard from '../components/Feed/PostCard';
import LeftSidebar from '../components/Profile/LeftSidebar';
import RightSidebar from '../components/Community/RightSidebar';
import { getPosts } from '../services/api'; // Import service API

const HomePage = () => {
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  // Buat fungsi untuk fetch data agar bisa dipanggil ulang
  const fetchPosts = useCallback(async () => {
    try {
      setLoading(true);
      const fetchedPosts = await getPosts();
      setPosts(fetchedPosts || []); // Pastikan posts adalah array
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  }, []);

  // Panggil fetchPosts saat komponen pertama kali di-mount
  useEffect(() => {
    fetchPosts();
  }, [fetchPosts]);

  return (
    <div className="main-content">
      <LeftSidebar />

      <main className="feed">
        {/* Kirim fungsi fetchPosts sebagai prop agar CreatePost bisa memicu refresh */}
        <CreatePost onPostCreated={fetchPosts} />
        
        {loading && <p>Loading feed...</p>}
        {error && <p style={{ color: 'red' }}>Error: {error}</p>}
        
        {!loading && !error && posts.length === 0 && (
          <div className="card">
            <p>Belum ada postingan. Jadilah yang pertama!</p>
          </div>
        )}

        {!loading && !error && posts.map(post => (
          // Ubah 'author.name' menjadi 'authorName' sesuai response API
          <PostCard key={post.id} post={{
              ...post,
              author: { name: post.authorName, username: post.authorUsername },
              timestamp: new Date(post.createdAt).toLocaleString('id-ID')
          }} />
        ))}
      </main>
      
      <RightSidebar />
    </div>
  );
};

export default HomePage;