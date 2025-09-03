import React, { useState } from 'react';
import { FiHeart, FiMessageSquare, FiShare } from 'react-icons/fi';
import { useAuth } from '../../contexts/AuthContext';
import { likePost, unlikePost } from '../../services/api';
import './Feed.css';

const PostCard = ({ post }) => {
  const { author, timestamp, content, media } = post;
  const { isAuthenticated, token } = useAuth();

  // State untuk like, diinisialisasi dari data post jika ada
  const [isLiked, setIsLiked] = useState(post.isLiked || false);
  const [likeCount, setLikeCount] = useState(post.likeCount || 0);
  
  const handleLikeToggle = async () => {
    if (!isAuthenticated) {
      alert("Silakan login untuk menyukai postingan.");
      return;
    }

    // Optimistic update: langsung ubah UI
    const originalLiked = isLiked;
    const originalLikeCount = likeCount;
    
    setIsLiked(!isLiked);
    setLikeCount(isLiked ? likeCount - 1 : likeCount + 1);

    try {
      if (originalLiked) {
        await unlikePost(post.id, token);
      } else {
        await likePost(post.id, token);
      }
    } catch (error) {
      console.error("Failed to toggle like:", error);
      // Jika error, kembalikan UI ke state semula
      setIsLiked(originalLiked);
      setLikeCount(originalLikeCount);
    }
  };

  return (
    <div className="card post-card">
      <div className="post-header">
        <img 
          src={`https://api.dicebear.com/8.x/initials/svg?seed=${author.name}`} 
          alt="Author Avatar" 
          className="avatar" 
        />
        <div className="post-author-info">
          <span className="author-name">{author.name}</span>
          <span className="post-timestamp">{timestamp}</span>
        </div>
        <div className="post-options">...</div>
      </div>
      
      <div className="post-content">
        <p>{content}</p>
        {media && <img src={media} alt="Post media" className="post-media" />}
      </div>

      <div className="post-actions">
        <button className={`action-button ${isLiked ? 'liked' : ''}`} onClick={handleLikeToggle}>
          <FiHeart /> <span>{likeCount > 0 ? likeCount : ''} Suka</span>
        </button>
        <button className="action-button">
          <FiMessageSquare /> <span>Komentar</span>
        </button>
        <button className="action-button">
          <FiShare /> <span>Bagikan</span>
        </button>
      </div>
    </div>
  );
};

export default PostCard;