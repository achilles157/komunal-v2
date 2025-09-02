import React from 'react';
import Navbar from './components/UI/Navbar/Navbar';
import './App.css';

// Placeholder untuk konten halaman
const HomePage = () => {
    // Di sini nanti kita akan meletakkan komponen-komponen lain
    // seperti LeftSidebar, Feed, dan RightSidebar
    return (
        <div className="main-content">
            <div className="left-sidebar">Left Sidebar</div>
            <div className="feed">Feed Area</div>
            <div className="right-sidebar">Right Sidebar</div>
        </div>
    );
}


function App() {
  return (
    <div className="App">
      <Navbar />
      <main className="app-content">
        <HomePage />
      </main>
    </div>
  );
}

export default App;