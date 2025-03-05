export const defaultFilters = {
    propType: 'player_goals',
    game: 'All Games',
    bookmaker: 'All'
};

export const propTypes = [
    { value: 'player_points', label: 'Points' },
    { value: 'player_power_play_points', label: 'Power Play Points' },
    { value: 'player_assists', label: 'Assists' },
    { value: 'player_goals', label: 'Goals' },
    { value: 'player_total_saves', label: 'Total Saves' },
];


export const bookmakers = [
    { value: 'All', label: 'All Bookmakers' },
    { value: 'betmgm', label: 'BetMGM' },
    { value: 'draftkings', label: 'DraftKings' },
    { value: 'fanduel', label: 'FanDuel' },
    { value: 'bovada', label: 'Bovada' },
    { value: 'betonlineag', label: 'BetOnline' },
    { value: 'williamhill_us', label: 'William Hill' },
    { value: 'betrivers', label: 'BetRivers' },
    { value: 'fanatics', label: 'Fanatics' },
];


export function getGameOptions(oddsData) {
    const options = ['All Games'];

    oddsData.forEach((game) => {

        const label = `${game.Away.Abbreviation} @ ${game.Home.Abbreviation}`;
        if (!options.includes(label)) {
            options.push(label);
        }
    });
    return options;
}

  export function  filterOddsData(oddsData, filters) {
    return oddsData
        .filter((game) => {
            if (filters.game === 'All Games') return true;
            const label = `${game.Away.Abbreviation} @ ${game.Home.Abbreviation}`;
            return filters.game === label;
        })
        .map((game) => {
            const homeTeamData = { ...game.Home };
            const awayTeamData = { ...game.Away };

            // Filter home roster
            if (homeTeamData.Roster) {
                homeTeamData.Roster = homeTeamData.Roster.filter((player) => {
                    if (!player.odds || !player.odds[filters.propType]) return false;
                    if (filters.bookmaker === 'All') {
                        return Object.keys(player.odds[filters.propType]).length > 0;
                    }
                    return !!player.odds[filters.propType][filters.bookmaker];
                });
            }

            // Filter away roster
            if (awayTeamData.Roster) {
                awayTeamData.Roster = awayTeamData.Roster.filter((player) => {
                    if (!player.odds || !player.odds[filters.propType]) return false;
                    if (filters.bookmaker === 'All') {
                        return Object.keys(player.odds[filters.propType]).length > 0;
                    }
                    return !!player.odds[filters.propType][filters.bookmaker];
                });
            }

            return {
                ...game,
                HomeTeam: homeTeamData,
                AwayTeam: awayTeamData
            };
        });
}
