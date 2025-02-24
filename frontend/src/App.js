// App.js
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import HomePage from './components/HomePage';
import GameDetails from './components/GameDetails';
import PlayerDetails from './components/PlayerDetails';
import './App.css';

function App() {
    return (
        <Router>
            <div className="App">
                <Routes>
                    <Route path="/" element={<HomePage />} />
                    <Route path="/game/:gameId" element={<GameDetails />} />
                    <Route path="/player/:playerId" element={<PlayerDetails />} />
                </Routes>
            </div>
        </Router>
    );
}

export default App;
