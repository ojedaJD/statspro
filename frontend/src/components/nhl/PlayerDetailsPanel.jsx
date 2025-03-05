import React, { useState, useMemo, useEffect } from 'react';
import {
    Box,
    Typography,
    Button,
    Paper,
    Chip,
    Divider,
    FormControl,
    InputLabel,
    Select,
    MenuItem,
    List,
    ListItem,
    ListItemText
} from '@mui/material';
import {
    BarChart,
    Bar,
    XAxis,
    YAxis,
    CartesianGrid,
    ReferenceLine,
    Tooltip,
    ResponsiveContainer,
    Cell
} from 'recharts';
import axios from "axios";

// Hockey-specific prop categories and mapping
const goaliePropCategories = ['Assists', 'Total Saves'];

const playerPropsCategories = ['Points', 'Goals', 'Assists', 'Power Play Points'];

const propTypeMapping = {
    'Points': 'player_points',
    'Goals': 'player_goals',
    'Assists': 'player_assists',
    'Power Play Points': 'player_power_play_points',
    'Total Saves': 'player_total_saves'
};

const reversePropsMapping = Object.entries(propTypeMapping).reduce((acc, [key, value]) => {
    acc[value] = key;
    return acc;
}, {});

// Derive available bookmakers dynamically from player's odds for the selected prop type
// (Defaults to "All" if none found)
const getAvailableBookmakers = (player, selectedApiPropType) => {
    if (player?.odds && player.odds[selectedApiPropType]) {
        return ['All', ...Object.keys(player.odds[selectedApiPropType])];
    }
    return ['All'];
};

// Helper function to get hockey logs (goalie or skater)
function getPlayerLogs(player) {
    return player.GoalieLogs ?? player.GameLogs ?? [];
}

const PlayerDetailView = ({ player, onBack, initialFilters = {} }) => {
    // For hockey we use hockey-specific prop categories.
    const initialProp = reversePropsMapping[initialFilters.propType] || 'Goals';
    const [selectedProp, setSelectedProp] = useState(initialProp);
    const [selectedBookmaker, setSelectedBookmaker] = useState(initialFilters.bookmaker || 'All');

    // Use hockey logs from either GoalieLogs or GameLogs
    const [gameLogs, setGameLogs] = useState(getPlayerLogs(player));

    // State to store the parsed odds data for the selected prop type
    const [propLines, setPropLines] = useState([]);

    const selectedApiPropType = propTypeMapping[selectedProp];

    // Update propLines whenever player data or selection changes
    useEffect(() => {
        if (player?.odds && player.odds[selectedApiPropType]) {
            const allOdds = [];
            const oddsObj = player.odds[selectedApiPropType];

            if (selectedBookmaker === 'All') {
                Object.entries(oddsObj).forEach(([bookmaker, bookmakerOdds]) => {
                    bookmakerOdds.forEach(odd => {
                        allOdds.push({
                            bookmaker,
                            ...odd
                        });
                    });
                });
            } else if (oddsObj[selectedBookmaker]) {
                oddsObj[selectedBookmaker].forEach(odd => {
                    allOdds.push({
                        bookmaker: selectedBookmaker,
                        ...odd
                    });
                });
            }
            setPropLines(allOdds);
        } else {
            setPropLines([]);
        }
    }, [player, selectedApiPropType, selectedBookmaker]);

    // Update available bookmakers dynamically based on odds data
    const availableBookmakers = useMemo(() => getAvailableBookmakers(player, selectedApiPropType), [player, selectedApiPropType]);

    const handlePropChange = (prop) => {
        setSelectedProp(prop);
    };

    const handleBookmakerChange = (event) => {
        setSelectedBookmaker(event.target.value);
    };

// \    'Points': 'player_points',
//         'Goals': 'player_goals',
//         'Assists': 'player_assists',
//         'Power Play Points': 'player_power_play_points',
//         'Total Saves': 'player_total_saves'

    const getValueFromLog = (log, prop) => {
        switch (prop) {
            case 'Total Saves':
                // Ensure fields exist; you might need to adjust if your data structure differs.
                return (log.shotsAgainst ?? 0) - (log.goalsAgainst ?? 0);
            case 'Goals':
                return log.goals || 0;
            case 'Points':
                return log.points || 0;
            case 'Power Play Points':
                return log.powerPlayPoints || 0;
            case 'Assists':
                return log.assists || 0;
            default:
                return 0;
        }
    };

    // Build chart data from gameLogs + selectedProp
    const chartData = useMemo(() => {
        if (!gameLogs || !Array.isArray(gameLogs)) return [];

        return gameLogs.map((game) => {
            return {
                opponent: game.opponentAbbrev || '',
                gameDate: game.gameDate ? game.gameDate.slice(5, 10) : '',
                value: getValueFromLog(game, selectedProp)
            };
        });
    }, [gameLogs, selectedProp]);

    // Get the most common prop line value across all bookmakers
    const getCurrentLine = useMemo(() => {
        if (!propLines.length) return null;
        const pointCounts = propLines.reduce((acc, odd) => {
            acc[odd.point] = (acc[odd.point] || 0) + 1;
            return acc;
        }, {});
        const [mostCommonPoint] = Object.entries(pointCounts).sort((a, b) => b[1] - a[1]);
        return mostCommonPoint ? mostCommonPoint[0] : null;
    }, [propLines]);

    // Decide color for each bar (green if value is at or above the line, red if below)
    const getBarColor = (value) => {
        if (!getCurrentLine) return '#8884d8';
        return value >= getCurrentLine ? '#4caf50' : '#f44336';
    };

    // Get best odds for over and under for the current line
    const getOddsDetails = useMemo(() => {
        if (!propLines.length || !getCurrentLine) {
            return { overOdds: '-', underOdds: '-' };
        }
        const matchingOdds = propLines.filter(odd => odd.point.toString() === getCurrentLine.toString());
        const overOdds = matchingOdds
            .filter(odd => odd.name === 'Over')
            .map(odd => odd.price)
            .sort((a, b) => b - a)[0] || '-';

        const underOdds = matchingOdds
            .filter(odd => odd.name === 'Under')
            .map(odd => odd.price)
            .sort((a, b) => b - a)[0] || '-';

        return { overOdds, underOdds };
    }, [propLines, getCurrentLine]);

    // For hockey, display the player's name differently since the data keys are different.
    const playerName = player?.firstName?.default
        ? `${player.firstName.default} ${player.lastName?.default || ''}`
        : (player?.name || 'Unknown Player');
    const teamAbbrev = player?.teamAbbrev || player?.TEAM_ABBREVIATION || '';

    const propCategories = player.positionCode === "G" ? goaliePropCategories : playerPropsCategories

    return (
        <Box sx={{ p: 2 }}>
            {/* Back Button */}
            <Button onClick={onBack} variant="text" sx={{ mb: 1, color: 'primary.main' }}>
                &lt; Back to Trends
            </Button>

            {/* Player Name / Team */}
            <Typography variant="h5" sx={{ fontWeight: 'bold', mb: 2 }}>
                {playerName} {teamAbbrev && `(${teamAbbrev})`}
            </Typography>

            {/* Prop Category Selection */}
            <Box sx={{ display: 'flex', gap: 2, flexWrap: 'wrap', mb: 2 }}>
                <Box>
                    <Typography variant="subtitle2" sx={{ mb: 1 }}>
                        Select a Prop:
                    </Typography>
                    <Box sx={{ display: 'flex', gap: 1, flexWrap: 'wrap' }}>
                        {propCategories.map((prop) => (
                            <Chip
                                key={prop}
                                label={prop}
                                onClick={() => handlePropChange(prop)}
                                color={selectedProp === prop ? 'primary' : 'default'}
                            />
                        ))}
                    </Box>
                </Box>

                {/* Bookmaker Selection */}
                <Box sx={{ minWidth: 200 }}>
                    <Typography variant="subtitle2" sx={{ mb: 1 }}>
                        Bookmaker:
                    </Typography>
                    <FormControl fullWidth size="small">
                        <Select value={selectedBookmaker} onChange={handleBookmakerChange} displayEmpty>
                            {availableBookmakers.map((bookmaker) => (
                                <MenuItem key={bookmaker} value={bookmaker}>
                                    {bookmaker}
                                </MenuItem>
                            ))}
                        </Select>
                    </FormControl>
                </Box>
            </Box>

            <Divider sx={{ mb: 2 }} />

            {/* Current Line & Odds */}
            <Box sx={{ display: 'flex', gap: 3, mb: 3, alignItems: 'center' }}>
                <Box>
                    <Typography variant="caption" color="text.secondary">
                        Current Line
                    </Typography>
                    <Typography variant="h6" sx={{ fontWeight: 'bold' }}>
                        {getCurrentLine || '-'}
                    </Typography>
                </Box>
                <Box>
                    <Typography variant="caption" color="text.secondary">
                        Over
                    </Typography>
                    <Typography variant="body1" color="success.main">
                        {getOddsDetails.overOdds > 0 ? `+${getOddsDetails.overOdds}` : getOddsDetails.overOdds}
                    </Typography>
                </Box>
                <Box>
                    <Typography variant="caption" color="text.secondary">
                        Under
                    </Typography>
                    <Typography variant="body1" color="error.main">
                        {getOddsDetails.underOdds > 0 ? `+${getOddsDetails.underOdds}` : getOddsDetails.underOdds}
                    </Typography>
                </Box>
            </Box>

            {/* Available Lines by Bookmaker */}
            <Paper sx={{ p: 2, mb: 3 }} elevation={2}>
                <Typography variant="subtitle1" sx={{ mb: 1, fontWeight: 'bold' }}>
                    {selectedProp} Lines by Bookmaker
                </Typography>
                {propLines.length > 0 ? (
                    <List dense>
                        {Object.entries(
                            propLines.reduce((acc, line) => {
                                if (!acc[line.point]) acc[line.point] = {};
                                if (!acc[line.point][line.bookmaker]) {
                                    acc[line.point][line.bookmaker] = [];
                                }
                                acc[line.point][line.bookmaker].push(line);
                                return acc;
                            }, {})
                        )
                            .sort((a, b) => parseFloat(a[0]) - parseFloat(b[0]))
                            .map(([point, bookmakers]) => (
                                <Box key={point} sx={{ mb: 2 }}>
                                    <Typography variant="body2" sx={{ fontWeight: 'bold' }}>
                                        Line: {point}
                                    </Typography>
                                    <Box sx={{ pl: 2 }}>
                                        {Object.entries(bookmakers).map(([bookmaker, lines]) => {
                                            const sortedLines = [...lines].sort((a, b) => {
                                                if (a.name === 'Over') return -1;
                                                if (b.name === 'Over') return 1;
                                                return 0;
                                            });
                                            return (
                                                <Box key={bookmaker} sx={{ display: 'flex', gap: 2 }}>
                                                    <Typography variant="body2" sx={{ minWidth: 100 }}>
                                                        {bookmaker}:
                                                    </Typography>
                                                    <Box sx={{ display: 'flex', gap: 2 }}>
                                                        {sortedLines.map((line, idx) => (
                                                            <Typography
                                                                key={idx}
                                                                variant="body2"
                                                                color={line.name === 'Over' ? 'success.main' : 'error.main'}
                                                            >
                                                                {line.name} {line.price > 0 ? '+' : ''}{line.price}
                                                            </Typography>
                                                        ))}
                                                    </Box>
                                                </Box>
                                            );
                                        })}
                                    </Box>
                                </Box>
                            ))}
                    </List>
                ) : (
                    <Typography variant="body2" color="text.secondary">
                        No odds data available for this prop.
                    </Typography>
                )}
            </Paper>

            {/* Bar Chart for Game Logs */}
            <Paper sx={{ p: 2 }} elevation={3}>
                <Box sx={{ width: '100%', height: 300 }}>
                    <ResponsiveContainer width="100%" height="100%">
                        <BarChart data={chartData}>
                            <CartesianGrid strokeDasharray="3 3" />
                            <XAxis dataKey="opponent" tick={{ fontSize: 10 }} />
                            <YAxis />
                            <Tooltip />
                            {getCurrentLine && (
                                <ReferenceLine y={getCurrentLine} stroke="black" strokeDasharray="3 3" />
                            )}
                            <Bar dataKey="value" label={{ position: 'top', fill: '#444' }} fill="#8884d8">
                                {chartData.map((entry, index) => (
                                    <Cell key={`cell-${index}`} fill={getBarColor(entry.value)} />
                                ))}
                            </Bar>
                        </BarChart>
                    </ResponsiveContainer>
                </Box>
            </Paper>
        </Box>
    );
};

export default PlayerDetailView;
