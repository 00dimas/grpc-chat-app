// src/App.tsx
import React from 'react';
import './App.css';
import Chat from './components/Chat.tsx'; // Mengimpor komponen Chat

const App: React.FC = () => {
  return (
    <div className="App">
      <h1>Welcome to the Chat App</h1>
      <Chat /> {/* Menambahkan komponen Chat di sini */}
    </div>
  );
};

export default App;
