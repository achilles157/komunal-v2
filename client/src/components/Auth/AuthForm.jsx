import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom'; // Import useNavigate
import { useAuth } from '../../contexts/AuthContext'; // Import useAuth
import './AuthForm.css';

const AuthForm = ({ isRegister = false }) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [fullName, setFullName] = useState('');
  const [username, setUsername] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

  const navigate = useNavigate(); // Hook untuk navigasi
  const { login } = useAuth(); // Ambil fungsi login dari context

  const handleSubmit = async (event) => {
    event.preventDefault();
    setLoading(true);
    setError('');

    const endpoint = isRegister ? '/api/register' : '/api/login';
    const payload = isRegister 
      ? { fullName, username, email, password } 
      : { email, password };

    try {
      const response = await fetch(endpoint, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(payload),
      });

      const data = await response.json();

      if (!response.ok) {
        throw new Error(data.message || 'Something went wrong');
      }

      if (isRegister) {
        alert('Registration successful! Please log in.');
        navigate('/login'); // Gunakan navigate untuk pindah halaman
      } else {
        // Gunakan fungsi login dari context!
        login(data.token);
        alert('Login successful!');
        navigate('/'); // Redirect ke halaman utama
      }

    } catch (err) {
      setError(err.message || 'An unknown error occurred');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="auth-container">
      <h2>{isRegister ? 'Buat Akun Baru' : 'Selamat Datang Kembali'}</h2>
      <p>{isRegister ? 'Bergabung dengan komunitas kreatif.' : 'Masuk untuk melanjutkan.'}</p>
      <form onSubmit={handleSubmit} className="auth-form">
        {isRegister && (
          <>
            <div className="form-group">
              <label htmlFor="fullName">Nama Lengkap</label>
              <input type="text" id="fullName" value={fullName} onChange={(e) => setFullName(e.target.value)} required />
            </div>
            <div className="form-group">
              <label htmlFor="username">Username</label>
              <input type="text" id="username" value={username} onChange={(e) => setUsername(e.target.value)} required />
            </div>
          </>
        )}
        <div className="form-group">
          <label htmlFor="email">Email</label>
          <input type="email" id="email" value={email} onChange={(e) => setEmail(e.target.value)} required />
        </div>
        <div className="form-group">
          <label htmlFor="password">Password</label>
          <input type="password" id="password" value={password} onChange={(e) => setPassword(e.target.value)} required />
        </div>
        
        {error && <p className="error-message">{error}</p>}

        <button type="submit" className="auth-button" disabled={loading}>
          {loading ? 'Memproses...' : (isRegister ? 'Daftar' : 'Masuk')}
        </button>
      </form>
      <div className="auth-switch">
        {isRegister ? 'Sudah punya akun? ' : 'Belum punya akun? '}
        <a href={isRegister ? '/login' : '/register'}>
          {isRegister ? 'Masuk di sini' : 'Daftar di sini'}
        </a>
      </div>
    </div>
  );
};

export default AuthForm;