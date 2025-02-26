// App.js
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import HomePage from './components/HomePage';
import PlayerDetails from './components/PlayerDetails';
import './App.css';
import NbaPage from "./components/nba/NbaPage";

function App() {
    return (
        <Router>
            <div className="App">
                <Routes>
                    <Route path="/" element={<HomePage />} />
                    <Route path="/nbapage" element={<NbaPage />} />
                    <Route path="/nba/" element={<PlayerDetails />} />
                </Routes>
            </div>
        </Router>
    );
}

export default App;
