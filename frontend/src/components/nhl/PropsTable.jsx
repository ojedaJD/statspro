import React, { useMemo, useState } from 'react';
import Box from '@mui/material/Box';
import Paper from '@mui/material/Paper';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell, { tableCellClasses } from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import TableSortLabel from '@mui/material/TableSortLabel';
import Typography from '@mui/material/Typography';
import Chip from '@mui/material/Chip';
import MenuItem from '@mui/material/MenuItem';
import FormControl from '@mui/material/FormControl';
import InputLabel from '@mui/material/InputLabel';
import Select from '@mui/material/Select';
import { styled } from '@mui/material/styles';

import { defaultFilters, getGameOptions, filterOddsData, propTypes, bookmakers } from '../../utils/nhlfilters';

// ---- NEW CODE: We can define a custom style or use inline style for coloring table cells.
const ColoredTableCell = styled(TableCell)(({ theme, colorbg }) => ({
    backgroundColor: colorbg || 'inherit',
    // You can do more styling if you like.
}));

/** ==========================
 *   TABLE FILTER COMPONENT
 * ========================== */
const TableFilterBar = ({ filters, onFilterChange, gameOptions }) => {
    return (
        <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 2, mb: 2 }}>
            <FormControl sx={{ minWidth: 200 }}>
                <InputLabel id="table-prop-type-label">Prop Type</InputLabel>
                <Select
                    labelId="table-prop-type-label"
                    id="table-prop-type-select"
                    value={filters.propType}
                    label="Prop Type"
                    onChange={(e) => onFilterChange('propType', e.target.value)}
                >
                    {propTypes.map((type) => (
                        <MenuItem key={type.value} value={type.value}>
                            {type.label}
                        </MenuItem>
                    ))}
                </Select>
            </FormControl>

            <FormControl sx={{ minWidth: 200 }}>
                <InputLabel id="table-game-label">Game</InputLabel>
                <Select
                    labelId="table-game-label"
                    id="table-game-select"
                    value={filters.game}
                    label="Game"
                    onChange={(e) => onFilterChange('game', e.target.value)}
                >
                    {gameOptions.map((option) => (
                        <MenuItem key={option} value={option}>
                            {option}
                        </MenuItem>
                    ))}
                </Select>
            </FormControl>

            <FormControl sx={{ minWidth: 200 }}>
                <InputLabel id="table-bookmaker-label">Bookmaker</InputLabel>
                <Select
                    labelId="table-bookmaker-label"
                    id="table-bookmaker-select"
                    value={filters.bookmaker}
                    label="Bookmaker"
                    onChange={(e) => onFilterChange('bookmaker', e.target.value)}
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

/** ==============================
 *   HELPER FUNCTIONS
 *  ============================== */

//
// 'player_points': 'Points',
//     'player_power_play_points': 'Power Play Points',
//     'player_assists': 'Assists',
//     'player_goals': 'Goals',
//     'player_total_saves': 'Total Saves',
//
// Extract which box‐score field to check, given the propType
function getBoxScoreValueFromLog(gameLog, propType) {
    if (gameLog)
    switch (propType) {
        case 'player_points':
            return gameLog.points;
        case 'player_power_play_points':
            return gameLog.powerPlayPoints;
        case 'player_goals':
            return gameLog.goals;
        case 'player_assists':
            return gameLog.assists;
        case 'player_total_saves':
            return gameLog.shotsAgainst - gameLog.goalsAgainst;

        default:
            return 0
    }
}

// Return the most common “line” used in the odds
function getMostCommonPoint(row) {
    const allOdds = Object.values(row.odds).flat();
    if (!allOdds.length) return '-';
    const pointCounts = allOdds.reduce((acc, odd) => {
        acc[odd.point] = (acc[odd.point] || 0) + 1;
        return acc;
    }, {});
    const [ mostCommon ] = Object.entries(pointCounts).sort((a,b) => b[1]-a[1]);
    return mostCommon ? parseFloat(mostCommon[0]) : '-';
}

/**
 * Compute how many of the player's logs are "hits" for the chosen line.
 *  - logs: an array of game logs
 *  - line: the numeric threshold (e.g. 24.5)
 *  - propType: which stat to measure
 *  - sliceN: optionally, only look at last N logs (if we want L20, L10, L5)
 * Returns a decimal fraction of hits from 0..1
 */
function calcHitFraction(logs, line, propType, sliceN=null) {
    const numericLine = parseFloat(line);
    if (isNaN(numericLine)) return 0;



    // Sort logs by game date if needed or assume they're in order already?
    // We'll assume the logs are in chronological order from newest to oldest or vice versa.
    // If you need the newest 5, you might do logs.slice(-5) etc.
    let relevantLogs = logs;

    if (sliceN) {
        relevantLogs = relevantLogs.slice(0,sliceN);



        // If logs are sorted oldest->newest, slicing from the end gets the newest N games
        // You might need to confirm your logs order or .reverse() if needed
    }

    if (!relevantLogs.length) return 0;

    let hits = 0;
    relevantLogs.forEach(g => {
        const val = getBoxScoreValueFromLog(g, propType);
        if (val >= numericLine) hits++;
    });
    return hits / relevantLogs.length;
}

/**
 * Compute "streak" of consecutive hits from the player's most recent logs
 * (We assume logs are from oldest => newest or vice versa; you need to confirm the order).
 * We'll assume the last entry is the newest game. We start from the end and count consecutive hits.
 */
function calcHitStreak(logs, line, propType) {
    const numericLine = parseFloat(line);
    if (isNaN(numericLine)) return 0;

    // We'll assume logs are in chronological order from oldest to newest,
    // so the newest game is logs[logs.length - 1].
    let streak = 0;
    for (let i =0; i <= logs.length; i++) {
        const val = getBoxScoreValueFromLog(logs[i], propType);
        if (val >= numericLine) {
            streak++;
        } else {
            break;
        }
    }
    return streak;
}

// Return the single line for the row
function getPropLine(row) {
    const line = getMostCommonPoint(row);
    return line === '-' ? '-' : line;
}

// Return best Over price
function getBestOver(row) {
    const line = getMostCommonPoint(row);
    if (line === '-') return '-';
    const allOdds = Object.values(row.odds).flat();
    const overOdds = allOdds
        .filter(o => o.name === 'Over' && o.point === line)
        .map(o => o.price);
    return overOdds.length ? Math.max(...overOdds) : '-';
}

// Return best Under price
function getBestUnder(row) {
    const line = getMostCommonPoint(row);
    if (line === '-') return '-';
    const allOdds = Object.values(row.odds).flat();
    const underOdds = allOdds
        .filter(o => o.name === 'Under' && o.point === line)
        .map(o => o.price);
    return underOdds.length ? Math.max(...underOdds) : '-';
}

// Example 1: use an HSL gradient from red to green for 0..100% rates.
export function getRateColor(rate) {
    // Pastel range: hue 0→120, fixed saturation ~45%, lightness ~85%.
    const hue = (120 * rate) / 100;
    // Feel free to adjust saturation/lightness for your taste:
    const saturation = 45;
    const lightness = 85;
    return `hsl(${hue}, ${saturation}%, ${lightness}%)`;
}

export function getStreakColor(streak) {
    // Cap the streak at some "max"—if your data has a typical ceiling,
    // you might pick 8 or 10 or whatever. Higher streaks all get the same
    // color at the upper limit.
    const maxStreak = 8;
    const clamped = Math.min(streak, maxStreak);

    // We'll map 0..maxStreak onto hue = 0..120,
    // same saturation/lightness as getRateColor for a consistent look.
    const hue = (120 * clamped) / maxStreak;
    const saturation = 45;
    const lightness = 85;
    return `hsl(${hue}, ${saturation}%, ${lightness}%)`;
}

// Build out the rows from our odds data
function getPlayerRows(oddsData, propType, bookMaker) {
    const rows = [];
    oddsData.forEach(game => {
        // home
        if (game.HomeTeam?.Roster) {
            game.HomeTeam.Roster.forEach(player => {
                if (player.odds && player.odds[propType]) {
                    rows.push({
                        id: player.id,
                        player,
                        team: game.Home.Abbreviation,
                        odds: bookMaker === 'All'
                            ? player.odds[propType]
                            : { [bookMaker]: player.odds[propType][bookMaker] || [] },
                        game: `${game.Away.Abbreviation} @ ${game.Home.Abbreviation}`
                    });
                }
            });
        }

        // away
        if (game.AwayTeam?.Roster) {
            game.AwayTeam.Roster.forEach(player => {
                if (player.odds && player.odds[propType]) {
                    rows.push({
                        id: player.id,
                        player,
                        team: game.Away.Abbreviation,
                        odds: bookMaker === 'All'
                            ? player.odds[propType]
                            : { [bookMaker]: player.odds[propType][bookMaker] || [] },
                        game: `${game.Away.Abbreviation} @ ${game.Home.Abbreviation}`
                    });
                }
            });
        }
    });
    return rows;
}

function getPlayerLogs(player) {
    // If a player is a goalie, they'll typically have "GoalieLogs".
    // If a player is a skater, they'll have "GameLogs".
    // You can adapt this if your data structure is different.
    return player.GoalieLogs ?? player.GameLogs ?? [];
}


/** ==========================
 *   MAIN COMPONENT
 * ========================== */
const PropsTable = ({ oddsData, onPlayerSelect }) => {
    const [order, setOrder] = useState('asc');
    const [orderBy, setOrderBy] = useState('name');
    const [tableFilters, setTableFilters] = useState(defaultFilters);

    const handleFilterChange = (filterType, value) => {
        setTableFilters((prev) => {
            const newFilters = { ...prev, [filterType]: value };
            return newFilters;
        });
    };


    const gameOptions = useMemo(() => getGameOptions(oddsData), [oddsData]);
    const filteredOddsData = useMemo(() => filterOddsData(oddsData, tableFilters), [oddsData, tableFilters]);

    const handleRequestSort = (property) => {
        const isAsc = orderBy === property && order === 'asc';
        setOrder(isAsc ? 'desc' : 'asc');
        setOrderBy(property);
    };

    const createSortHandler = (property) => () => handleRequestSort(property);

    const rows = useMemo(() => {
        return getPlayerRows(filteredOddsData, tableFilters.propType, tableFilters.bookmaker);
    }, [filteredOddsData, tableFilters.propType, tableFilters.bookmaker]);

    // Sorting
    const sortedRows = useMemo(() => {
        const comparator = (a, b) => {
            let aValue, bValue;
            switch (orderBy) {
                case 'player':
                    aValue = `${a.player.firstName.default} ${a.player.lastName.default}`;
                    bValue = `${b.player.firstName.default} ${b.player.lastName.default}`;
                    break;
                case 'team':
                    aValue = a.team;
                    bValue = b.team;
                    break;
                case 'position':
                    aValue = a.player.positionCode;
                    bValue = b.player.positionCode;
                    break;
                case 'line':
                    aValue = parseFloat(getPropLine(a)) || 0;
                    bValue = parseFloat(getPropLine(b)) || 0;
                    break;
                case 'over':
                    aValue = parseFloat(getBestOver(a)) || 0;
                    bValue = parseFloat(getBestOver(b)) || 0;
                    break;
                case 'under':
                    aValue = parseFloat(getBestUnder(a)) || 0;
                    bValue = parseFloat(getBestUnder(b)) || 0;
                    break;

                // Season, L20, L10, L5 => We'll compute them on the fly
                case 'hitSeason': {
                    const lineA = parseFloat(getPropLine(a)) || 0;
                    const lineB = parseFloat(getPropLine(b)) || 0;
                    aValue = calcHitFraction(getPlayerLogs(a.player), lineA, tableFilters.propType) * 100;
                    bValue = calcHitFraction(getPlayerLogs(b.player), lineB, tableFilters.propType) * 100;
                    break;
                }
                case 'hitL20': {
                    const lineA = parseFloat(getPropLine(a)) || 0;
                    const lineB = parseFloat(getPropLine(b)) || 0;
                    aValue = calcHitFraction(getPlayerLogs(a.player), lineA, tableFilters.propType, 20)*100;
                    bValue = calcHitFraction(getPlayerLogs(b.player), lineB, tableFilters.propType, 20)*100;
                    break;
                }
                case 'hitL10': {
                    const lineA = parseFloat(getPropLine(a)) || 0;
                    const lineB = parseFloat(getPropLine(b)) || 0;
                    aValue = calcHitFraction(getPlayerLogs(a.player), lineA, tableFilters.propType, 10)*100;
                    bValue = calcHitFraction(getPlayerLogs(b.player), lineB, tableFilters.propType, 10)*100;
                    break;
                }
                case 'hitL5': {
                    const lineA = parseFloat(getPropLine(a)) || 0;
                    const lineB = parseFloat(getPropLine(b)) || 0;
                    aValue = calcHitFraction(getPlayerLogs(a.player), lineA, tableFilters.propType, 5)*100;
                    bValue = calcHitFraction(getPlayerLogs(b.player), lineB, tableFilters.propType, 5)*100;
                    break;
                }
                case 'streak': {
                    const lineA = parseFloat(getPropLine(a)) || 0;
                    const lineB = parseFloat(getPropLine(b)) || 0;
                    aValue = calcHitStreak(getPlayerLogs(a.player), lineA, tableFilters.propType);
                    bValue = calcHitStreak(getPlayerLogs(b.player), lineB, tableFilters.propType);
                    break;
                }
                default:
                    aValue = a[orderBy];
                    bValue = b[orderBy];
                    break;
            }

            if (order === 'asc') {
                return aValue < bValue ? -1 : aValue > bValue ? 1 : 0;
            } else {
                return aValue > bValue ? -1 : aValue < bValue ? 1 : 0;
            }
        };

        return [...rows].sort(comparator);
    }, [rows, order, orderBy, tableFilters.propType]);

    // Label for the chosen prop
    const getPropTypeLabel = () => {
        const propTypeMap = {
            'player_points': 'Points',
            'player_power_play_points': 'Power Play Points',
            'player_assists': 'Assists',
            'player_goals': 'Goals',
            'player_total_saves': 'Total Saves',
        };
        return propTypeMap[tableFilters.propType] || tableFilters.propType;
    };

    return (
        <Paper elevation={2} sx={{ overflowY: 'auto', maxHeight: 'calc(100vh - 10vh)' }}>
            <Box sx={{ p: 2, borderBottom: '1px solid rgba(224, 224, 224, 1)' }}>
                <Typography variant="h6" sx={{ mb: 2 }}>
                    {getPropTypeLabel()} Props
                </Typography>
                <TableFilterBar
                    filters={tableFilters}
                    onFilterChange={handleFilterChange}
                    gameOptions={gameOptions}
                />
            </Box>

            <TableContainer>
                <Table sx={{ minWidth: 1000 }} aria-label="props table" align={"center"}>
                    <TableHead>
                        <TableRow>
                            <TableCell align={"center"}>
                                <TableSortLabel
                                    active={orderBy === 'player'}
                                    direction={orderBy === 'player' ? order : 'asc'}
                                    onClick={createSortHandler('player')}
                                >
                                    First
                                </TableSortLabel>
                            </TableCell>
                            <TableCell align={"center"}>
                                <TableSortLabel
                                    active={orderBy === 'team'}
                                    direction={orderBy === 'team' ? order : 'asc'}
                                    onClick={createSortHandler('team')}
                                >
                                    Team
                                </TableSortLabel>
                            </TableCell>
                            <TableCell align={"center"}>
                                <TableSortLabel
                                    active={orderBy === 'position'}
                                    direction={orderBy === 'position' ? order : 'asc'}
                                    onClick={createSortHandler('position')}
                                >
                                    Position
                                </TableSortLabel>
                            </TableCell>

                            <TableCell align={"center"}>
                                <TableSortLabel
                                    active={orderBy === 'line'}
                                    direction={orderBy === 'line' ? order : 'asc'}
                                    onClick={createSortHandler('line')}
                                >
                                    Line
                                </TableSortLabel>
                            </TableCell>
                            <TableCell align={"center"}>
                                <TableSortLabel
                                    active={orderBy === 'over'}
                                    direction={orderBy === 'over' ? order : 'asc'}
                                    onClick={createSortHandler('over')}
                                >
                                    Best Over
                                </TableSortLabel>
                            </TableCell>
                            <TableCell align={"center"}>
                                <TableSortLabel
                                    active={orderBy === 'under'}
                                    direction={orderBy === 'under' ? order : 'asc'}
                                    onClick={createSortHandler('under')}
                                >
                                    Best Under
                                </TableSortLabel>
                            </TableCell>

                            {/** NEW COLUMNS: Season / L20 / L10 / L5 / Streak */}
                            <TableCell align={"center"}>
                                <TableSortLabel
                                    active={orderBy === 'hitSeason'}
                                    direction={orderBy === 'hitSeason' ? order : 'asc'}
                                    onClick={createSortHandler('hitSeason')}
                                >
                                    Season
                                </TableSortLabel>
                            </TableCell>
                            <TableCell align={"center"}>
                                <TableSortLabel
                                    active={orderBy === 'hitL20'}
                                    direction={orderBy === 'hitL20' ? order : 'asc'}
                                    onClick={createSortHandler('hitL20')}
                                >
                                    L20
                                </TableSortLabel>
                            </TableCell>
                            <TableCell align={"center"}>
                                <TableSortLabel
                                    active={orderBy === 'hitL10'}
                                    direction={orderBy === 'hitL10' ? order : 'asc'}
                                    onClick={createSortHandler('hitL10')}
                                >
                                    L10
                                </TableSortLabel>
                            </TableCell>
                            <TableCell align={"center"}>
                                <TableSortLabel
                                    active={orderBy === 'hitL5'}
                                    direction={orderBy === 'hitL5' ? order : 'asc'}
                                    onClick={createSortHandler('hitL5')}
                                >
                                    L5
                                </TableSortLabel>
                            </TableCell>
                            <TableCell align={"center"}>
                                <TableSortLabel
                                    active={orderBy === 'streak'}
                                    direction={orderBy === 'streak' ? order : 'asc'}
                                    onClick={createSortHandler('streak')}
                                >
                                    Streak
                                </TableSortLabel>
                            </TableCell>
                        </TableRow>
                    </TableHead>

                    <TableBody>
                        {sortedRows.length === 0 ? (
                            <TableRow>
                                <TableCell colSpan={10} align="center">
                                    <Typography variant="body2" sx={{ py: 2 }}>
                                        No props data available.
                                    </Typography>
                                </TableCell>
                            </TableRow>
                        ) : (
                            sortedRows.map((row,index) => {
                                const line = getPropLine(row);
                                const bestOver = getBestOver(row);
                                const bestUnder = getBestUnder(row);
                                const logs = getPlayerLogs(row.player);
                                const position = row.player?.positionCode

                                const seasonFrac = calcHitFraction(logs, line, tableFilters.propType, null) * 100;
                                const l20Frac = calcHitFraction(logs, line, tableFilters.propType, 20) * 100;
                                const l10Frac = calcHitFraction(logs, line, tableFilters.propType, 10) * 100;
                                const l5Frac  = calcHitFraction(logs, line, tableFilters.propType, 5)  * 100;

                                const streakVal = calcHitStreak(logs, line, tableFilters.propType);

                                return (
                                    <TableRow
                                        hover
                                        key={`${row.id}-${index}`}
                                        onClick={() => onPlayerSelect(row.player,tableFilters)}
                                        sx={{ cursor: 'pointer', '&:hover': { backgroundColor: 'action.hover' } }}
                                    >
                                        <TableCell align={"center"}>
                                            <Typography variant="body2" sx={{ fontWeight: 'bold' }}>
                                                {row.player.firstName.default} {row.player.lastName.default}
                                            </Typography>
                                        </TableCell>
                                        <TableCell align={"center"}>
                                            <Typography variant="body2">{row.team}</Typography>
                                        </TableCell>
                                        <TableCell align={"center"}>
                                            <Typography variant="body2">{position}</Typography>
                                        </TableCell>

                                        <TableCell align={"center"}>
                                            <Typography variant="body2" sx={{ fontWeight: 'medium' }}>
                                                {line}
                                            </Typography>
                                        </TableCell>
                                        <TableCell align={"center"}>
                                            <Chip
                                                label={bestOver}
                                                size="small"
                                                color={bestOver > 0 ? 'success' : 'default'}
                                                sx={{ fontWeight: 'bold', minWidth: 60 }}
                                            />
                                        </TableCell>
                                        <TableCell align={"center"}>
                                            <Chip
                                                label={bestUnder}
                                                size="small"
                                                color={bestUnder > 0 ? 'success' : 'default'}
                                                sx={{ fontWeight: 'bold', minWidth: 60 }}
                                            />
                                        </TableCell>

                                        {/** NEW: Season, L20, L10, L5 with background color */}
                                        <ColoredTableCell colorbg={getRateColor(seasonFrac)} align={"center"}>
                                            <Typography variant="body2" sx={{ fontWeight: 'bold' }}>
                                                {logs.length === 0 ? '-' : `${Math.round(seasonFrac)}%`}
                                            </Typography>
                                        </ColoredTableCell>

                                        <ColoredTableCell colorbg={getRateColor(l20Frac)} align={"center"}>
                                            <Typography variant="body2" sx={{ fontWeight: 'bold' }}>
                                                {logs.length === 0 ? '-' : `${Math.round(l20Frac)}%`}
                                            </Typography>
                                        </ColoredTableCell>

                                        <ColoredTableCell colorbg={getRateColor(l10Frac)} align={"center"}>
                                            <Typography variant="body2" sx={{ fontWeight: 'bold' }}>
                                                {logs.length === 0 ? '-' : `${Math.round(l10Frac)}%`}
                                            </Typography>
                                        </ColoredTableCell>

                                        <ColoredTableCell colorbg={getRateColor(l5Frac)} align={"center"}>
                                            <Typography variant="body2" sx={{ fontWeight: 'bold' }}>
                                                {logs.length === 0 ? '-' : `${Math.round(l5Frac)}%`}
                                            </Typography>
                                        </ColoredTableCell>

                                        {/** Streak also color-coded */}
                                        <ColoredTableCell colorbg={getStreakColor(streakVal)} align={"center"}>
                                            <Typography variant="body2" sx={{ fontWeight: 'bold' }}>
                                                {streakVal}
                                            </Typography>
                                        </ColoredTableCell>
                                    </TableRow>
                                );
                            })
                        )}
                    </TableBody>
                </Table>
            </TableContainer>
        </Paper>
    );
};

export default PropsTable;
