// components/nba/PlayerPanel.jsx
import React from "react";

const PlayerPanel = ({ playerName, imageUrl }) => {
    return (
        <div className="player-panel">
            <h2>{playerName}</h2>
            <img src={imageUrl} alt={`${playerName}`} />
        </div>
    );
};

export default PlayerPanel;
