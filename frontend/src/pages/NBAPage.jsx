import React, { useState, useEffect } from 'react';
import { Routes, Route, useNavigate } from 'react-router-dom';
import axios from 'axios';
import Box from '@mui/material/Box';
import Grid from '@mui/material/Grid2'; // Using Grid2 as specified
import Typography from '@mui/material/Typography';
import CircularProgress from '@mui/material/CircularProgress';
import Alert from '@mui/material/Alert';
import Sidebar from '../components/Sidebar';
import PropsTable from '../components/PropsTable';
import PlayerDetailView from '../components/PlayerDetailsPanel';
import { defaultFilters } from '../utils/filters';




const NBAPage = () => {
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [oddsData, setOddsData] = useState([]);
    const [selectedPlayer, setSelectedPlayer] = useState(null);
    const [playerGameLog, setPlayerGameLog] = useState([]);
    const [showPlayerPanel, setShowPlayerPanel] = useState(false);
    const [globalFilters, setGlobalFilters] = useState(defaultFilters);


    const updateGlobalFilters = (newFilters) => {
        setGlobalFilters((prev) => ({ ...prev, ...newFilters }));
    };

    // Build a list of game labels dynamically
    const gameOptions = ['All Games'];
    if (oddsData)
    oddsData.forEach(game => {
        const gameLabel = `${game.AwayTeam.abbreviation} @ ${game.HomeTeam.abbreviation}`;
        if (!gameOptions.includes(gameLabel)) {
            gameOptions.push(gameLabel);
        }
    });


    const navigate = useNavigate();

    useEffect(() => {
        const fetchOddsData = async () => {
            try {
                setLoading(true);
                const response = await axios.get('http://localhost:8080/nba/v2/matchups');
                setOddsData(response.data);
                setError(null);
            } catch (err) {
                console.error('Error fetching odds data:', err);
                setError('Failed to load props data. Please try again later.');
            } finally {
                setLoading(false);
            }
        };

        fetchOddsData();
    }, []);

    const fetchPlayerGameLog = async (playerId) => {
        try {
            setLoading(true);
            const response = await axios.get(
                `http://localhost:8080/nba/v2/player/gamelogs?playerID=${playerId}&season=2024-25&seasonType=Regular+Season`
            );
            setPlayerGameLog(response.data);
        } catch (err) {
            console.error('Error fetching player game log:', err);
            setError('Failed to load player data. Please try again later.');
        } finally {
            setLoading(false);
        }
    };

    const handlePlayerSelect = (player) => {
        setSelectedPlayer(player);

        setShowPlayerPanel(true);
    };

    const handleClosePlayerPanel = () => {
        setShowPlayerPanel(false);
    };


    return (
   <>
            <Grid container spacing={2}>
                <Grid size={{ xs: 12, md: 3 , lg: 2 }}>
                    <Sidebar
                        oddsData={oddsData}
                        onPlayerSelect={(player) => handlePlayerSelect(player)}
                        updateGlobalFilters={updateGlobalFilters}
                    />
                </Grid>

                <Grid size={{ xs: 12, md: 9, lg: 10 }}>
                    {loading && <CircularProgress sx={{ display: 'block', margin: '40px auto' }} />}

                    {error && (
                        <Alert severity="error" sx={{ sm: 2 }}>
                            {error}
                        </Alert>
                    )}

                    {!loading && !error && (
                        showPlayerPanel ? (
                            <PlayerDetailView
                                player={selectedPlayer}
                                onBack={() => setShowPlayerPanel(false)}
                                initialFilters={globalFilters}
                            />
                        ) : (
                            <PropsTable
                                oddsData={oddsData}
                                onPlayerSelect={(player) => handlePlayerSelect(player)}
                                updateGlobalFilters={updateGlobalFilters}

                            />
                        )
                    )}
                </Grid>


            </Grid>

</>
    );
};

export default NBAPage;