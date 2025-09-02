import React from 'react';
import AuthForm from '../components/Auth/AuthForm';

const LoginPage = () => {
  return (
    <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '80vh' }}>
      <AuthForm isRegister={false} />
    </div>
  );
};

export default LoginPage;