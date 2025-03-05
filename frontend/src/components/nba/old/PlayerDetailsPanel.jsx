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

// Recharts imports
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

const propCategories = [
    'Pts',
    'Reb',
    'Ast',
    'PA',
    'PR',
    'RA',
    'PRA',
    '3PTM',
    'BLK',
    'STL',
    'TOV',
    'STL + BLK'
];

const splitFilters = [
    'Season',
    '1H',
    '2H',
    '1Q',
    '2Q',
    '3Q',
    '4Q'
];

// Map from propCategories to the API prop types
const propTypeMapping = {
    'Pts': 'player_points',
    'Reb': 'player_rebounds',
    'Ast': 'player_assists',
    'PA': 'player_points_assists',
    'PR': 'player_points_rebounds',
    'RA': 'player_rebounds_assists',
    'PRA': 'player_points_rebounds_assists',
    '3PTM': 'player_threes',
    'BLK': 'player_blocks',
    'STL': 'player_steals',
    'TOV': 'player_turnovers',
    'STL + BLK': 'player_blocks_steals'
};

// Reverse mapping for displaying in UI
const reversePropsMapping = Object.entries(propTypeMapping).reduce((acc, [key, value]) => {
    acc[value] = key;
    return acc;
}, {});

const bookmakers = [
    { value: 'All', label: 'All Bookmakers' },
    { value: 'draftkings', label: 'DraftKings' },
    { value: 'fanduel', label: 'FanDuel' },
];

const PlayerDetailView = ({ player, onBack, initialFilters = {} }) => {

    const initialPropType = reversePropsMapping[initialFilters.propType] || 'Pts';
    const [selectedProp, setSelectedProp] = useState(initialPropType);
    const [selectedSplit, setSelectedSplit] = useState('Season');
    const [selectedBookmaker, setSelectedBookmaker] = useState(initialFilters.bookmaker || 'All');
    const [gameLogs, setGameLogs] = useState(player?.CurrentSeasonLogs || []);

    // State to store the parsed odds data for the selected prop type
    const [propLines, setPropLines] = useState([]);

    const selectedApiPropType = propTypeMapping[selectedProp];

    // Extract all bookmaker odds when prop type changes or player changes
    useEffect(() => {
        if (player?.odds && player.odds[selectedApiPropType]) {
            const allOdds = [];
            const oddsObj = player.odds[selectedApiPropType];

            // Extract all odds for the selected bookmaker or all bookmakers
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

    const handlePropChange = (prop) => {
        setSelectedProp(prop);
    };

    const handleBookmakerChange = (event) => {
        setSelectedBookmaker(event.target.value);
    };

    const handleSplitChange = async (split) => {
        setSelectedSplit(split);

        try {
            const response = await axios.get(`http://localhost:8080/nba/v2/player/gamelogs`, {
                params: {
                    playerID: player.PERSON_ID,
                    period: split // Ensure it matches backend expected format
                }
            });

            setGameLogs(response.data);
        } catch (error) {
            console.error('Error fetching split game logs:', error);
        }
    };

    // Get the current line to display in the chart
    const getCurrentLine = useMemo(() => {
        if (!propLines.length) return null;

        // Get the most common point value across all bookmakers
        const pointCounts = propLines.reduce((acc, odd) => {
            acc[odd.point] = (acc[odd.point] || 0) + 1;
            return acc;
        }, {});

        const [mostCommonPoint] = Object.entries(pointCounts)
            .sort((a, b) => b[1] - a[1]);


        return mostCommonPoint ? mostCommonPoint[0] : null;
    }, [propLines]);

    // Decide which numeric stat from the game log to chart
    const getValueFromLog = (log, prop) => {
        switch (prop) {
            case 'Pts':       return log.PTS;
            case 'Reb':       return log.REB;
            case 'Ast':       return log.AST;
            case 'PA':        return (log.PTS + log.AST);
            case 'PR':        return (log.PTS + log.REB);
            case 'RA':        return (log.REB + log.AST);
            case 'PRA':       return (log.PTS + log.REB + log.AST);
            case '3PTM':      return log.FG3M;
            case 'BLK':       return log.BLK;
            case 'STL':       return log.STL;
            case 'TOV':       return log.TOV;
            case 'STL + BLK': return (log.STL + log.BLK);
            default:          return log.PTS; // fallback
        }
    };

    // Build chart data from gameLogs + selectedProp
    const chartData = useMemo(() => {
        if (!gameLogs || !Array.isArray(gameLogs)) return [];

        return gameLogs.map((game) => {
            const value = getValueFromLog(game, selectedProp);

            return {
                opponent: game.MATCHUP.substring(3,11)|| '',
                gameDate: game.GAME_DATE?.slice(5, 10) || '', // e.g. "02-26"
                value
            };
        });
    }, [gameLogs, selectedProp]);

    // Decide color for each bar (green if over line, red if under)
    const getBarColor = (value) => {
        if (!getCurrentLine) return '#8884d8'; // Default color if no line
        return value >= getCurrentLine ? '#4caf50' : '#f44336'; // green : red
    };

    // Get odds details for the Over/Under display
    const getOddsDetails = useMemo(() => {
        if (!propLines.length || !getCurrentLine) {
            return { overOdds: '-', underOdds: '-' };
        }

        // Find the best odds for over and under for the current line
        const matchingOdds = propLines.filter(odd => odd.point.toString() === getCurrentLine.toString());
        // Get best over odds
        const overOdds = matchingOdds
            .filter(odd => odd.name === 'Over')
            .map(odd => odd.price)
            .sort((a, b) => b - a)[0] || '-';


        // Get best under odds
        const underOdds = matchingOdds
            .filter(odd => odd.name === 'Under')
            .map(odd => odd.price)
            .sort((a, b) => b - a)[0] || '-';

        return { overOdds, underOdds };
    }, [propLines, getCurrentLine]);

    return (
        <Box sx={{ p: 2 }}>
            {/* Back to Trends */}
            <Button
                onClick={onBack}
                variant="text"
                sx={{ mb: 1, color: 'primary.main' }}
            >
                &lt; Back to Trends
            </Button>

            {/* Player Name / Team */}
            <Typography variant="h5" sx={{ fontWeight: 'bold', mb: 2 }}>
                {player?.DISPLAY_FIRST_LAST} ({player?.TEAM_ABBREVIATION})
            </Typography>

            {/* Prop Category and Splits */}
            <Box sx={{ display: 'flex', gap: 2, flexWrap: 'wrap', mb: 2 }}>
                {/* Prop categories */}
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

                <Box>
                    <Typography variant="subtitle2" sx={{ mb: 1 }}>
                        Game Splits:
                    </Typography>
                    <Box sx={{ display: 'flex', gap: 1, flexWrap: 'wrap' }}>
                        {splitFilters.map((split) => (
                            <Chip
                                key={split}
                                label={split}
                                onClick={() => handleSplitChange(split)}
                                color={selectedSplit === split ? 'primary' : 'default'}
                            />
                        ))}
                    </Box>
                </Box>

                {/* Bookmaker selection */}
                <Box sx={{ minWidth: 200 }}>
                    <Typography variant="subtitle2" sx={{ mb: 1 }}>
                        Bookmaker:
                    </Typography>
                    <FormControl fullWidth size="small">
                        <Select
                            value={selectedBookmaker}
                            onChange={handleBookmakerChange}
                            displayEmpty
                        >
                            {bookmakers.map((bookmaker) => (
                                <MenuItem key={bookmaker.value} value={bookmaker.value}>
                                    {bookmaker.label}
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
                        {getOddsDetails.overOdds > 0
                            ? `+${getOddsDetails.overOdds}`
                            : getOddsDetails.overOdds}
                    </Typography>
                </Box>
                <Box>
                    <Typography variant="caption" color="text.secondary">
                        Under
                    </Typography>
                    <Typography variant="body1" color="error.main">
                        {getOddsDetails.underOdds > 0
                            ? `+${getOddsDetails.underOdds}`
                            : getOddsDetails.underOdds}
                    </Typography>
                </Box>
            </Box>

            {/* Available Lines from Different Bookmakers */}
            <Paper sx={{ p: 2, mb: 3 }} elevation={2}>
                <Typography variant="subtitle1" sx={{ mb: 1, fontWeight: 'bold' }}>
                    {selectedProp} Lines by Bookmaker
                </Typography>
                {propLines.length > 0 ? (
                    <List dense>
                        {/* Group by point, then by bookmaker */}
                        {Object.entries(
                            propLines.reduce((acc, line) => {
                                if (!acc[line.point]) acc[line.point] = {};
                                if (!acc[line.point][line.bookmaker]) {
                                    acc[line.point][line.bookmaker] = [];
                                }
                                acc[line.point][line.bookmaker].push(line);
                                return acc;
                            }, {})
                        ).sort((a, b) => parseFloat(a[0]) - parseFloat(b[0])).map(([point, bookmakers]) => (
                            <Box key={point} sx={{ mb: 2 }}>
                                <Typography variant="body2" sx={{ fontWeight: 'bold' }}>
                                    Line: {point}
                                </Typography>
                                <Box sx={{ pl: 2 }}>
                                    {Object.entries(bookmakers).map(([bookmaker, lines]) => {
                                        // Sort lines to ensure "Over" comes before "Under"
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

            {/* Bar Chart of Game Logs */}
            <Paper sx={{ p: 2 }} elevation={3}>
                <Box sx={{ width: '100%', height: 300 }}>
                    <ResponsiveContainer width="100%" height="100%">
                        <BarChart data={chartData}>
                            <CartesianGrid strokeDasharray="3 3" />
                            <XAxis dataKey="opponent" tick={{ fontSize: 10 }} />
                            <YAxis />
                            <Tooltip />
                            {/* The reference line to show the current over/under line */}
                            {getCurrentLine && (
                                <ReferenceLine y={getCurrentLine} stroke="black" strokeDasharray="3 3" />
                            )}
                            <Bar
                                dataKey="value"
                                label={{ position: 'top', fill: '#444' }}
                                fill="#8884d8"
                            >
                                {
                                    // Use the "children" function pattern to style each bar individually
                                    chartData.map((entry, index) => (
                                        <Cell
                                            key={`cell-${index}`}
                                            fill={getBarColor(entry.value)}
                                        />
                                    ))
                                }
                            </Bar>
                        </BarChart>
                    </ResponsiveContainer>
                </Box>
            </Paper>
        </Box>
    );
};

export default PlayerDetailView;