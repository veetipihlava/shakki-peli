import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Home from './components/Home.tsx';
import GameRoom from './components/GameRoom.tsx';
import './App.css';

const App: React.FC = () => {
  return (
    <Router>
      <div className="container">
        <h1>Multiplayer Game</h1>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/game/:gameID" element={<GameRoom />} />
        </Routes>
      </div>
    </Router>
  );
};

export default App;