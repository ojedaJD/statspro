import React from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import { ThemeProvider, createTheme } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import Layout from './components/Layout';
import NBAPage from './pages/NBAPage';
import UnderConstruction from './pages/UnderConstruction';

// Create a theme with primary and secondary colors
const theme = createTheme({
    palette: {
        primary: {
            main: '#2e7d34', // Green color similar to props.cash
        },
        secondary: {
            main: '#f50057',
        },
        background: {
            default: '#f5f5f0',
        },
    },
    typography: {
        fontFamily: '"Roboto", "Helvetica", "Arial", sans-serif',
    },
    components: {
        MuiGrid2: {
            defaultProps: {
                // Set default spacing for Grid2
                spacing: 2,
            },
        },
    },
});

function App() {
    return (
        <ThemeProvider theme={theme}>
            <CssBaseline />
            <Router>
                <Routes>
                    <Route path="/" element={<Layout />}>
                        <Route index element={<Navigate to="/nba" replace />} />
                        <Route path="nba/*" element={<NBAPage />} />
                        <Route path="nfl" element={<UnderConstruction sport="NFL" />} />
                        <Route path="nhl" element={<UnderConstruction sport="NHL" />} />
                        <Route path="mlb" element={<UnderConstruction sport="MLB" />} />
                        <Route path="ncaaf" element={<UnderConstruction sport="NCAAF" />} />
                        <Route path="wnba" element={<UnderConstruction sport="WNBA" />} />
                    </Route>
                </Routes>
            </Router>
        </ThemeProvider>
    );
}

export default App;