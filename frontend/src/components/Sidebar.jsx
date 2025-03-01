import React, { useState, useMemo } from 'react';
import Box from '@mui/material/Box';
import Paper from '@mui/material/Paper';
import Typography from '@mui/material/Typography';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import Divider from '@mui/material/Divider';
import Chip from '@mui/material/Chip';
import FormControl from '@mui/material/FormControl';
import InputLabel from '@mui/material/InputLabel';
import Select from '@mui/material/Select';
import MenuItem from '@mui/material/MenuItem';
import { defaultFilters, getGameOptions, filterOddsData, propTypes, bookmakers } from '../utils/filters';
import {all} from "axios";


// Portrait-oriented filter component for Sidebar
const SidebarFilterBar = ({ filters, onFilterChange, gameOptions }) => {
    return (
        <Box sx={{ display: 'flex', flexDirection: 'column', gap: 2, mb: 2 }}>
            <FormControl fullWidth>
                <InputLabel id="sidebar-prop-type-label">Prop Type</InputLabel>
                <Select
                    labelId="sidebar-prop-type-label"
                    id="sidebar-prop-type-select"
                    value={filters.propType}
                    label="Prop Type"
                    onChange={(e) => onFilterChange('propType', e.target.value)}
                    size="small"
                >
                    {propTypes.map((type) => (
                        <MenuItem key={type.value} value={type.value}>
                            {type.label}
                        </MenuItem>
                    ))}
                </Select>
            </FormControl>

            <FormControl fullWidth>
                <InputLabel id="sidebar-game-label">Game</InputLabel>
                <Select
                    labelId="sidebar-game-label"
                    id="sidebar-game-select"
                    value={filters.game}
                    label="Game"
                    onChange={(e) => onFilterChange('game', e.target.value)}
                    size="small"
                >
                    {gameOptions.map((option) => (
                        <MenuItem key={option} value={option}>
                            {option}
                        </MenuItem>
                    ))}
                </Select>
            </FormControl>

            <FormControl fullWidth>
                <InputLabel id="sidebar-bookmaker-label">Bookmaker</InputLabel>
                <Select
                    labelId="sidebar-bookmaker-label"
                    id="sidebar-bookmaker-select"
                    value={filters.bookmaker}
                    label="Bookmaker"
                    onChange={(e) => onFilterChange('bookmaker', e.target.value)}
                    size="small"
                >
                    {bookmakers.map((bookmaker) => (
                        <MenuItem key={bookmaker.value} value={bookmaker.value}>
                            {bookmaker.label}
                        </MenuItem>
                    ))}
                </Select>
            </FormControl>
        </Box>
    );
};




const Sidebar = ({ oddsData, onPlayerSelect, updateGlobalFilters }) => {
    // Internal state for sidebar filters
    const [sidebarFilters, setSidebarFilters] = useState(defaultFilters);

    const handleFilterChange = (filterType, value) => {
        setSidebarFilters((prev) => {
            const newFilters = { ...prev, [filterType]: value };
            updateGlobalFilters({ [filterType]: value }); // Only updates global reference
            return newFilters;
        });
    };

    const gameOptions = useMemo(() => getGameOptions(oddsData), [oddsData]);

    const filteredOddsData = useMemo(() => {
        return filterOddsData(oddsData, sidebarFilters);
    }, [oddsData, sidebarFilters]);

    console.log(filteredOddsData)

    // Get prop type label for display
    const getPropTypeLabel = () => {
        const propTypeMap = {
            'player_points': 'Points',
            'player_rebounds': 'Rebounds',
            'player_assists': 'Assists',
            'player_points_rebounds_assists': 'PRA',
            'player_points_rebounds': 'Points + Rebounds',
            'player_points_assists': 'Points + Assists',
            'player_steals': 'Steals',
            'player_blocks': 'Blocks',
            'player_turnovers': 'Turnovers',
            'player_threes': '3PM'
        };
        return propTypeMap[sidebarFilters.propType] || sidebarFilters.propType;
    };

    const renderPlayerLine = (player) => {
        const propType = sidebarFilters.propType;
        const bookMaker = sidebarFilters.bookmaker;

        if (!player.odds || !player.odds[propType]) {
            return null;
        }

        let allOdds = [];

        // Get all odds based on bookmaker filter
        if (bookMaker === "All") {
            // Collect odds from all bookmakers
            Object.values(player.odds[propType]).forEach(bookmakerOdds => {
                allOdds = [...allOdds, ...bookmakerOdds];
            });
        } else if (player.odds[propType][bookMaker]) {
            // Get odds for the specific bookmaker
            allOdds = player.odds[propType][bookMaker];
        } else {
            return null; // No odds for this bookmaker
        }

        if (allOdds.length === 0) {
            return null;
        }

        // **Step 1: Find the Most Common Point Value**
        const pointCounts = allOdds.reduce((acc, odd) => {
            acc[odd.point] = (acc[odd.point] || 0) + 1;
            return acc;
        }, {});

        const mostCommonPoint = Object.keys(pointCounts)
            .map(Number)
            .sort((a, b) => pointCounts[b] - pointCounts[a])[0] || '-';

        // **Step 2: Find the Best Odds for Over & Under**
        const overOdds = allOdds
            .filter(odd => odd.name === 'Over' && odd.point === mostCommonPoint)
            .map(odd => odd.price)
            .sort((a, b) => b - a)[0] || '-';

        const underOdds = allOdds
            .filter(odd => odd.name === 'Under' && odd.point === mostCommonPoint)
            .map(odd => odd.price)
            .sort((a, b) => b - a)[0] || '-';

        return (
            <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', width: '100%' }}>
                <Box sx={{ flex: 1 }}>
                    <Typography variant="subtitle2" noWrap>
                        {mostCommonPoint}
                    </Typography>
                </Box>
                <Box sx={{ display: 'flex', gap: 1 }}>
                    <Chip
                        label={`O ${overOdds > 0 ? '+' + overOdds : overOdds}`}
                        size="small"
                        color={overOdds > 0 ? "success" : "default"}
                        sx={{
                            fontWeight: 'bold',
                            minWidth: 75,
                            fontSize: '0.75rem'
                        }}
                    />
                    <Chip
                        label={`U ${underOdds > 0 ? '+' + underOdds : underOdds}`}
                        size="small"
                        color={underOdds > 0 ? "success" : "default"}
                        sx={{
                            fontWeight: 'bold',
                            minWidth: 75,
                            fontSize: '0.75rem'
                        }}
                    />
                </Box>
            </Box>
        );
    };

    const renderTeamSection = (team, gameId) => {
        if (!team || !team.Roster) return null;

        const playersWithProps = team.Roster.filter(player =>
            player.odds && player.odds[sidebarFilters.propType]
        );

        if (playersWithProps.length === 0) return null;

        return (
            <Box sx={{ mb: 3 }}>
                <Typography variant="subtitle1" sx={{ fontWeight: 'bold', mb: 1 }}>
                    {team.abbreviation} {team.nickname}
                </Typography>
                <List disablePadding>
                    {playersWithProps.map((player) => (
                        <ListItem
                            key={player.PERSON_ID}
                            button
                            onClick={() => onPlayerSelect(player)}
                            sx={{
                                py: 1,
                                '&:hover': { bgcolor: 'action.hover' }
                            }}
                        >
                            <Box sx={{ width: '100%' }}>
                                <Box sx={{ display: 'flex', justifyContent: 'space-between', width: '100%', mb: 0.5 }}>
                                    <Typography variant="body2" sx={{ fontWeight: 'bold' }}>
                                        {player.DISPLAY_FIRST_LAST}
                                    </Typography>
                                    <Typography variant="caption" color="text.secondary">
                                        {team.abbreviation} - {player.ROSTERSTATUS === 1 ? 'Active' : 'Inactive'}
                                    </Typography>
                                </Box>
                                {renderPlayerLine(player)}
                            </Box>
                        </ListItem>
                    ))}
                </List>
            </Box>
        );
    };

    const renderGames = () => {
        return filteredOddsData.map((game, index) => {
            const homeTeamSection = renderTeamSection(game.HomeTeam, game.id);
            const awayTeamSection = renderTeamSection(game.AwayTeam, game.id);

            if (!homeTeamSection && !awayTeamSection) return null;

            const gameTitle = `${game.AwayTeam.abbreviation} @ ${game.HomeTeam.abbreviation}`;

            return (
                <Box key={index} sx={{ mb: 3 }}>
                    <Typography variant="h6" sx={{ mb: 1, fontWeight: 'bold' }}>
                        {gameTitle}
                    </Typography>
                    <Divider sx={{ mb: 2 }} />
                    {awayTeamSection}
                    {homeTeamSection}
                </Box>
            );
        });
    };

    return (
        <Paper
            elevation={2}
            sx={{
                p: 2,
                height: '100%',
                overflowY: 'auto',
                maxHeight: 'calc(100vh - 160px)',
            }}
        >
            <Typography variant="h6" sx={{ mb: 2 }}>
                {getPropTypeLabel()} Props
            </Typography>

            {/* Portrait-oriented filter component */}
            <SidebarFilterBar
                filters={sidebarFilters}
                onFilterChange={handleFilterChange}
                gameOptions={gameOptions}
            />

            <Divider sx={{ my: 2 }} />

            {filteredOddsData.length > 0 ? (
                renderGames()
            ) : (
                <Typography variant="body2" color="text.secondary">
                    No games available.
                </Typography>
            )}
        </Paper>
    );
};

export default Sidebar;