import React from 'react';
import { FiHeart, FiMessageSquare, FiShare } from 'react-icons/fi'; // Menggunakan ikon
import './Feed.css';

const PostCard = ({ post }) => {
  const { author, timestamp, content, media } = post;

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
        <button className="action-button">
          <FiHeart /> <span>Suka</span>
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