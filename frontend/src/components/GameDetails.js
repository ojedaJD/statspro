// components/GameDetails.js
import React from 'react';
import { useParams, Link } from 'react-router-dom';

const GameDetails = () => {
    const { gameId } = useParams();

    // Mock game data for demonstration
    const gameData = {
        id: gameId,
        players: [
            { id: 101, name: 'LeBron James', imageUrl: 'frontend/src/components/nba/assets/LebronHeadshot.jpg' },
            { id: 102, name: 'Stephen Curry', imageUrl: '/assets/CurryHeadshot.jpg' }
        ]
    };

    return (
        <div>
            <h1>Game Details for Game {gameId}</h1>
            <div className="players-list">
                {gameData.players.map(player => (
                    <div key={player.id} className="player-summary">
                        <Link to={`/player/${player.id}`}>
                            <img
                                src={player.imageUrl}
                                alt={player.name}
                                style={{ width: '150px', marginRight: '10px' }}
                            />
                            <p>{player.name}</p>
                        </Link>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default GameDetails;
