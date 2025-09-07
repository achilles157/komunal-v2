import React, { useState } from 'react';
import { useAuth } from '../contexts/AuthContext';
import { createCommunity } from '../services/api';
import { useNavigate } from 'react-router-dom';

const CreateCommunityPage = () => {
  const { token } = useAuth();
  const navigate = useNavigate();
  const [name, setName] = useState('');
  const [description, setDescription] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError('');
    try {
      const newCommunity = await createCommunity({ name, description }, token);
      alert(`Komunitas "${newCommunity.name}" berhasil dibuat!`);
      // Arahkan ke halaman komunitas baru menggunakan slug dari response API
      navigate(`/c/${newCommunity.slug}`); 
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="card" style={{ maxWidth: '700px', margin: '2rem auto' }}>
      <h2>Buat Komunitas Baru</h2>
      <p>Bangun ruang untuk para kreator dengan minat yang sama.</p>
      {error && <p style={{ color: 'red' }}>{error}</p>}
      <form onSubmit={handleSubmit} className="auth-form" style={{ textAlign: 'left' }}>
        <div className="form-group">
          <label htmlFor="name">Nama Komunitas</label>
          <input type="text" value={name} onChange={(e) => setName(e.target.value)} required />
        </div>
        <div className="form-group">
          <label htmlFor="description">Deskripsi</label>
          <textarea value={description} onChange={(e) => setDescription(e.target.value)} rows="4" style={{ resize: 'vertical', width: '100%', padding: '0.75rem', border: '1px solid #dbdbdb', borderRadius: '6px', fontSize: '1rem' }}></textarea>
        </div>
        <button type="submit" className="auth-button" disabled={loading}>
          {loading ? 'Membuat...' : 'Buat Komunitas'}
        </button>
      </form>
    </div>
  );
};

export default CreateCommunityPage;