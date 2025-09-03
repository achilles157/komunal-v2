import React, { useState, useEffect } from 'react';
import { useAuth } from '../contexts/AuthContext';
import { getUserProfile, updateUserProfile } from '../services/api';
import { useNavigate } from 'react-router-dom';

const EditProfilePage = () => {
  const { user, token } = useAuth();
  const navigate = useNavigate();
  
  const [formData, setFormData] = useState({
    fullName: '',
    profilePictureUrl: '',
    bio: '',
  });
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    if (user) {
      // Ambil data profil terbaru untuk diisi ke form
      getUserProfile(user.username)
        .then(data => {
          setFormData({
            fullName: data.fullName,
            profilePictureUrl: data.profilePictureUrl || '',
            bio: data.bio || '',
          });
          setLoading(false);
        })
        .catch(err => {
          setError(err.message);
          setLoading(false);
        });
    }
  }, [user]);

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError('');
    try {
      await updateUserProfile(formData, token);
      alert('Profil berhasil diperbarui!');
      navigate(`/${user.username}`); // Kembali ke halaman profil
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  if (!user) return <div className="app-content"><p>Silakan login untuk mengedit profil.</p></div>;

  return (
    <div className="card" style={{ maxWidth: '700px', margin: '2rem auto' }}>
      <h2>Edit Profil</h2>
      {error && <p style={{ color: 'red' }}>{error}</p>}
      <form onSubmit={handleSubmit} className="auth-form" style={{textAlign: 'left'}}>
        <div className="form-group">
          <label htmlFor="fullName">Nama Lengkap</label>
          <input type="text" name="fullName" value={formData.fullName} onChange={handleChange} required />
        </div>
        <div className="form-group">
          <label htmlFor="profilePictureUrl">URL Foto Profil</label>
          <input type="text" name="profilePictureUrl" value={formData.profilePictureUrl} onChange={handleChange} placeholder="https://example.com/image.png"/>
        </div>
        <div className="form-group">
          <label htmlFor="bio">Bio</label>
          <textarea name="bio" value={formData.bio} onChange={handleChange} rows="4" style={{resize: 'vertical', width: '100%', padding: '0.75rem', border: '1px solid #dbdbdb', borderRadius: '6px', fontSize: '1rem'}}></textarea>
        </div>
        <button type="submit" className="auth-button" disabled={loading}>
          {loading ? 'Menyimpan...' : 'Simpan Perubahan'}
        </button>
      </form>
    </div>
  );
};

export default EditProfilePage;