import React from 'react';
import AuthForm from '../components/Auth/AuthForm';

const RegisterPage = () => {
  return (
    <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '80vh' }}>
      <AuthForm isRegister={true} />
    </div>
  );
};

export default RegisterPage;