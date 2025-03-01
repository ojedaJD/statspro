import React from 'react';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import Box from '@mui/material/Box';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';

const PlayerDetail = ({ player, selectedBook, onBack }) => {
    // Check if the player has odds from the selected book
    const hasSelectedBookOdds = player.odds && player.odds[selectedBook];

    // Format player prop data for display
    const formatPropData = () => {
        if (!hasSelectedBookOdds) {
            return [];
        }

        return Object.entries(player.odds[selectedBook]).map(([propType, propData]) => {
            return {
                propType: formatPropType(propType),
                point: propData.point,
                name: propData.name,
                price: propData.price
            };
        }).sort((a, b) => a.propType.localeCompare(b.propType));
    };

    // Format the prop type for display
    const formatPropType = (propType) => {
        return propType
            .replace('player_', '')
            .replace('_', ' ')
            .split('_')
            .map(word => word.charAt(0).toUpperCase() + word.slice(1))
            .join(' + ');
    };

    const propData = formatPropData();

    return (
        <Box>
            <Box sx={{ display: 'flex', alignItems: 'center', mb: 3 }}>
                <Button
                    startIcon={<ArrowBackIcon />}
                    onClick={onBack}
                    variant="outlined"
                    sx={{ mr: 2 }}
                >
                    Back
                </Button>
                <Typography variant="h5" component="h2">
                    {player.name}
                </Typography>
            </Box>

            <Box sx={{ mb: 3 }}>
                <Typography variant="subtitle1" gutterBottom>
                    Team: {player.teamName} ({player.team})
                </Typography>
                <Typography variant="subtitle1" gutterBottom>
                    Matchup: {player.team} vs {player.gameInfo.opponent}
                </Typography>
                <Typography variant="subtitle1" gutterBottom>
                    Sportsbook: {selectedBook.charAt(0).toUpperCase() + selectedBook.slice(1).replace('_', ' ')}
                </Typography>
            </Box>

            {!hasSelectedBookOdds ? (
                <Typography variant="body1" color="text.secondary">
                    No odds available for this player from {selectedBook}.
                </Typography>
            ) : (
                <TableContainer component={Paper} variant="outlined">
                    <Table sx={{ minWidth: 650 }} size="small">
                        <TableHead>
                            <TableRow>
                                <TableCell>Prop Type</TableCell>
                                <TableCell>Line</TableCell>
                                <TableCell>Direction</TableCell>
                                <TableCell>Price</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {propData.map((prop, index) => (
                                <TableRow key={index}>
                                    <TableCell component="th" scope="row">
                                        {prop.propType}
                                    </TableCell>
                                    <TableCell>{prop.point}</TableCell>
                                    <TableCell>{prop.name}</TableCell>
                                    <TableCell>
                                        {prop.price > 0 ? `+${prop.price}` : prop.price}
                                    </TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
            )}
        </Box>
    );
};

export default PlayerDetail;