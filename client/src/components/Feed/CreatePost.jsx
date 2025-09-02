import React, { useState } from 'react';
import { useAuth } from '../../contexts/AuthContext';
import { createPost } from '../../services/api'; // Import service API
import './Feed.css';

const CreatePost = ({ onPostCreated }) => { // Terima prop onPostCreated
  const { isAuthenticated, user, token } = useAuth();
  const [postContent, setPostContent] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState('');

  if (!isAuthenticated) {
    return null;
  }

  const handlePost = async () => {
    setIsLoading(true);
    setError('');
    try {
      // Panggil API untuk membuat post baru
      await createPost({ content: postContent, mediaUrl: '' }, token);
      setPostContent(''); // Kosongkan input
      
      // Panggil fungsi callback dari parent untuk me-refresh feed
      if (onPostCreated) {
        onPostCreated();
      }

    } catch (err) {
      setError(err.message);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="card create-post-card">
      {error && <p style={{ color: 'red', textAlign: 'center' }}>{error}</p>}
      <div className="create-post-header">
        <img 
          src={`https://api.dicebear.com/8.x/initials/svg?seed=${user.username}`} 
          alt="User Avatar" 
          className="avatar" 
        />
        <textarea
          value={postContent}
          onChange={(e) => setPostContent(e.target.value)}
          placeholder={`Apa yang Anda pikirkan, ${user.username}?`}
          className="post-input"
          disabled={isLoading}
        />
      </div>
      <div className="create-post-actions">
        <button 
          onClick={handlePost} 
          className="post-button" 
          disabled={!postContent.trim() || isLoading}
        >
          {isLoading ? 'Memposting...' : 'Posting'}
        </button>
      </div>
    </div>
  );
};

export default CreatePost;