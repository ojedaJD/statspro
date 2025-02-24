// components/HomePage.js
import React from 'react';
import { Link } from 'react-router-dom';

const HomePage = () => {
    // Example game schedule data
    const gameSchedule = [
        { id: 1, teams: 'Lakers vs Warriors', time: '7:00 PM' },
    ];

    return (
        <div>
            <header>
                <h1>Current Game Schedule</h1>
                <div className="schedule-panel">
                    {gameSchedule.map(game => (
                        <div key={game.id} className="game-item">
                            <Link to={`/game/${game.id}`}>
                                {game.teams} - {game.time}
                            </Link>
                        </div>
                    ))}
                </div>
            </header>
        </div>
    );
};

export default HomePage;
