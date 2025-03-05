import React, { useState, useEffect } from 'react';
import { Routes, Route, useNavigate } from 'react-router-dom';
import axios from 'axios';
import Box from '@mui/material/Box';
import Grid from '@mui/material/Grid2'; // Using Grid2 as specified
import Typography from '@mui/material/Typography';
import CircularProgress from '@mui/material/CircularProgress';
import Alert from '@mui/material/Alert';
import Sidebar from '../components/nhl/Sidebar';
import PropsTable from '../components/nhl/PropsTable';
import PlayerDetailView from '../components/nhl/PlayerDetailsPanel';




const NHLPage = () => {
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [oddsData, setOddsData] = useState([]);
    const [selectedPlayer, setSelectedPlayer] = useState(null);
    const [showPlayerPanel, setShowPlayerPanel] = useState(false);
    const [selectedFilters, setSelectedFilters] = useState(null);


    // Build a list of game labels dynamically
    const gameOptions = ['All Games'];
    if (oddsData)
    oddsData.forEach(game => {
        const gameLabel = `${game.Away.Abbreviation} @ ${game.Home.Abbreviation}`;
        if (!gameOptions.includes(gameLabel)) {
            gameOptions.push(gameLabel);
        }
    });


    const navigate = useNavigate();

    useEffect(() => {
        const fetchOddsData = async () => {
            try {
                setLoading(true);
                const response = await axios.get('http://localhost:8080/nhl/matchups');
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

    const handlePlayerSelect = (player, filters) => {
        setSelectedPlayer(player);
        setSelectedFilters(filters);
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
                        onPlayerSelect={handlePlayerSelect}
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
                                initialFilters={selectedFilters}
                            />
                        ) : (
                            <PropsTable
                                oddsData={oddsData}
                                onPlayerSelect={handlePlayerSelect}
                            />
                        )
                    )}
                </Grid>


            </Grid>

</>
    );
};

export default NHLPage;