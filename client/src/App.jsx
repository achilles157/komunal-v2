import React from 'react';
import { Routes, Route } from 'react-router-dom'; // Hapus import BrowserRouter
import Navbar from './components/UI/Navbar/Navbar';
import HomePage from './pages/HomePage';
import LoginPage from './pages/LoginPage';
import RegisterPage from './pages/RegisterPage';
import ProfilePage from './pages/ProfilePage';
import EditProfilePage from './pages/EditProfilePage'; 
import CreateCommunityPage from './pages/CreateCommunityPage';
import CommunityPage from './pages/CommunityPage';
import './App.css';

function App() {
  return (
    <div className="App">
      <Navbar />
      <main className="app-content">
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/login" element={<LoginPage />} />
          <Route path="/register" element={<RegisterPage />} />
          <Route path="/:username" element={<ProfilePage />} />
          <Route path="/settings/profile" element={<EditProfilePage />} />
          <Route path="/new/community" element={<CreateCommunityPage />} />
          <Route path="/:username" element={<ProfilePage />} />
          <Route path="/c/:communityName" element={<CommunityPage />} /> 
          <Route path="/:username" element={<ProfilePage />} />
        </Routes>
      </main>
    </div>
  );
}

export default App;