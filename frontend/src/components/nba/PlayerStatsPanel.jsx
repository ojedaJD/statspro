import React, { useEffect, useState } from "react";
import axios from "axios";
import { Box, Button, Typography } from "@mui/material";

const PlayerStatsPanel = ({ player, onBack }) => {
    const [stats, setStats] = useState(null);
    const currentSeason = "2023-24";
    const seasonType = "Regular Season";

    useEffect(() => {
        if (!player) return;
        const fetchPlayerStats = async () => {
            try {
                const response = await axios.get(`http://localhost:8080/nba/player/gamelog`, {
                    params: {
                        playerID: player.id,
                        season: currentSeason,
                        seasonType: seasonType
                    }
                });
                setStats(response.data);
            } catch (error) {
                console.error("Error fetching player stats:", error);
            }
        };

        fetchPlayerStats();
    }, [player]);

    return (
        <Box sx={{ p: 3 }}>
            <Typography variant="h4">{player.DISPLAY_FIRST_LAST} - Stats</Typography>
            {stats ? (
                <Typography variant="body1">{JSON.stringify(stats, null, 2)}</Typography>
            ) : (
                <Typography variant="h6" sx={{ my: 2 }}>Loading stats...</Typography>
            )}
            <Button variant="contained" onClick={onBack}>Back to Player List</Button>
        </Box>
    );
};

export default PlayerStatsPanel;
