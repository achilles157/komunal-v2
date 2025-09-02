import React, { useState } from 'react';
import './AuthForm.css';

const AuthForm = ({ isRegister = false }) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [fullName, setFullName] = useState(''); // Hanya untuk register

  const handleSubmit = (event) => {
    event.preventDefault();
    const endpoint = isRegister ? '/api/register' : '/api/login';
    const payload = isRegister ? { fullName, email, password } : { email, password };
    
    console.log(`Mengirim data ke ${endpoint}`, payload);
    // TODO: Buat panggilan API ke backend menggunakan fetch atau axios
  };

  return (
    <div className="auth-container">
      <h2>{isRegister ? 'Buat Akun Baru' : 'Selamat Datang Kembali'}</h2>
      <p>{isRegister ? 'Bergabung dengan komunitas kreatif.' : 'Masuk untuk melanjutkan.'}</p>
      <form onSubmit={handleSubmit} className="auth-form">
        {isRegister && (
          <div className="form-group">
            <label htmlFor="fullName">Nama Lengkap</label>
            <input
              type="text"
              id="fullName"
              value={fullName}
              onChange={(e) => setFullName(e.target.value)}
              required
            />
          </div>
        )}
        <div className="form-group">
          <label htmlFor="email">Email</label>
          <input
            type="email"
            id="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>
        <div className="form-group">
          <label htmlFor="password">Password</label>
          <input
            type="password"
            id="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        <button type="submit" className="auth-button">
          {isRegister ? 'Daftar' : 'Masuk'}
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