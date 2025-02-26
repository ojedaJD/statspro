// components/PlayerDetails.js
import React from 'react';
import { useParams } from 'react-router-dom';
import PlayerPanel from './PlayerPanel';

const PlayerDetails = () => {
    const { playerId } = useParams();

    // Replace this mock data with a real data fetch if needed
    const playerData = {
        101: { name: 'LeBron James', imageUrl: 'frontend/src/components/nba/assets/LebronHeadshot.jpg' },
        102: { name: 'Stephen Curry', imageUrl: '/assets/CurryHeadshot.jpg' },
    };

    const player = playerData[playerId] || { name: 'Unknown Player', imageUrl: '' };

    return (
        <div className="player-details-page">
            <h1>Player Details</h1>
            <div className="stats-container">
                <div className="player-panel-container">
                    <PlayerPanel playerName={player.name} imageUrl={player.imageUrl} />
                </div>
                <div className="graph-container">
                    <h2>Player Stats Graph</h2>
                    <p>[Graph Placeholder]</p>
                </div>
            </div>
        </div>
    );
};

export default PlayerDetails;
