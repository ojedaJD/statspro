package nba

import (
	"encoding/json"
	"fmt"
	odds "sports_api/odds/nba"
	models "sports_api/stats/endpoints/nba"
	"strings"
)

type Matchup struct {
	HomeTeam *Team
	AwayTeam *Team
}
type Team struct {
	ID                int    `json:"id"`
	Abbreviation      string `json:"abbreviation"`
	Nickname          string `json:"nickname"`
	YearFounded       int    `json:"year_founded"`
	City              string `json:"city"`
	FullName          string `json:"full_name"`
	State             string `json:"state"`
	ChampionshipYears []int  `json:"championship_years"`
	Roster            []models.Player
}

func (t *Team) GetPlayerByName(name string) *models.Player {
	for i := range t.Roster {
		if t.Roster[i].Name == name {
			return &t.Roster[i]
		}
	}
	return nil
}

type Teams []Team

func (r *Team) addRosterMember(player models.Player) bool {
	if r == nil {
		return false
	}
	r.Roster = append(r.Roster, player)
	return true
}

// GetNBATeams returns a hardcoded list of NBA teams
func GetNBATeams() Teams {
	return []Team{
		{1610612737, "ATL", "Hawks", 1949, "Atlanta", "Atlanta Hawks", "Georgia", []int{1958}, []models.Player{}},
		{1610612738, "BOS", "Celtics", 1946, "Boston", "Boston Celtics", "Massachusetts", []int{1957, 1959, 1960, 1961, 1962, 1963, 1964, 1965, 1966, 1968, 1969, 1974, 1976, 1981, 1984, 1986, 2008}, []models.Player{}},
		{1610612739, "CLE", "Cavaliers", 1970, "Cleveland", "Cleveland Cavaliers", "Ohio", []int{2016}, []models.Player{}},
		{1610612740, "NOP", "Pelicans", 2002, "New Orleans", "New Orleans Pelicans", "Louisiana", []int{}, []models.Player{}},
		{1610612741, "CHI", "Bulls", 1966, "Chicago", "Chicago Bulls", "Illinois", []int{1991, 1992, 1993, 1996, 1997, 1998}, []models.Player{}},
		{1610612742, "DAL", "Mavericks", 1980, "Dallas", "Dallas Mavericks", "Texas", []int{2011}, []models.Player{}},
		{1610612743, "DEN", "Nuggets", 1976, "Denver", "Denver Nuggets", "Colorado", []int{2023}, []models.Player{}},
		{1610612744, "GSW", "Warriors", 1946, "Golden State", "Golden State Warriors", "California", []int{1947, 1956, 1975, 2015, 2017, 2018, 2022}, []models.Player{}},
		{1610612745, "HOU", "Rockets", 1967, "Houston", "Houston Rockets", "Texas", []int{1994, 1995}, []models.Player{}},
		{1610612746, "LAC", "Clippers", 1970, "Los Angeles", "Los Angeles Clippers", "California", []int{}, []models.Player{}},
		{1610612747, "LAL", "Lakers", 1948, "Los Angeles", "Los Angeles Lakers", "California", []int{1949, 1950, 1952, 1953, 1954, 1972, 1980, 1982, 1985, 1987, 1988, 2000, 2001, 2002, 2009, 2010, 2020}, []models.Player{}},
		{1610612748, "MIA", "Heat", 1988, "Miami", "Miami Heat", "Florida", []int{2006, 2012, 2013}, []models.Player{}},
		{1610612749, "MIL", "Bucks", 1968, "Milwaukee", "Milwaukee Bucks", "Wisconsin", []int{1971, 2021}, []models.Player{}},
		{1610612750, "MIN", "Timberwolves", 1989, "Minnesota", "Minnesota Timberwolves", "Minnesota", []int{}, []models.Player{}},
		{1610612751, "BKN", "Nets", 1976, "Brooklyn", "Brooklyn Nets", "New York", []int{}, []models.Player{}},
		{1610612752, "NYK", "Knicks", 1946, "New York", "New York Knicks", "New York", []int{1970, 1973}, []models.Player{}},
		{1610612753, "ORL", "Magic", 1989, "Orlando", "Orlando Magic", "Florida", []int{}, []models.Player{}},
		{1610612754, "IND", "Pacers", 1976, "Indiana", "Indiana Pacers", "Indiana", []int{}, []models.Player{}},
		{1610612755, "PHI", "76ers", 1949, "Philadelphia", "Philadelphia 76ers", "Pennsylvania", []int{1955, 1967, 1983}, []models.Player{}},
		{1610612756, "PHX", "Suns", 1968, "Phoenix", "Phoenix Suns", "Arizona", []int{}, []models.Player{}},
		{1610612757, "POR", "Trail Blazers", 1970, "Portland", "Portland Trail Blazers", "Oregon", []int{1977}, []models.Player{}},
		{1610612758, "SAC", "Kings", 1948, "Sacramento", "Sacramento Kings", "California", []int{1951}, []models.Player{}},
		{1610612759, "SAS", "Spurs", 1976, "San Antonio", "San Antonio Spurs", "Texas", []int{1999, 2003, 2005, 2007, 2014}, []models.Player{}},
		{1610612760, "OKC", "Thunder", 1967, "Oklahoma City", "Oklahoma City Thunder", "Oklahoma", []int{1979}, []models.Player{}},
		{1610612761, "TOR", "Raptors", 1995, "Toronto", "Toronto Raptors", "Ontario", []int{2019}, []models.Player{}},
		{1610612762, "UTA", "Jazz", 1974, "Utah", "Utah Jazz", "Utah", []int{}, []models.Player{}},
		{1610612763, "MEM", "Grizzlies", 1995, "Memphis", "Memphis Grizzlies", "Tennessee", []int{}, []models.Player{}},
		{1610612764, "WAS", "Wizards", 1961, "Washington", "Washington Wizards", "District of Columbia", []int{1978}, []models.Player{}},
		{1610612765, "DET", "Pistons", 1948, "Detroit", "Detroit Pistons", "Michigan", []int{1989, 1990, 2004}, []models.Player{}},
		{1610612766, "CHA", "Hornets", 1988, "Charlotte", "Charlotte Hornets", "North Carolina", []int{}, []models.Player{}},
	}
}
func (t Teams) GetTeamByID(teamID int) *Team {
	for i := range t { // Iterate by index to modify slice elements
		if t[i].ID == teamID {
			return &t[i] // Return the actual element in the slice, not a copy
		}
	}
	return nil
}

// GetWNBATeams returns a hardcoded list of WNBA teams
func GetWNBATeams() Teams {
	return []Team{
		{1611661313, "NYL", "Liberty", 1997, "New York", "New York Liberty", "New York", []int{}, []models.Player{}},
		{1611661317, "PHO", "Mercury", 1997, "Phoenix", "Phoenix Mercury", "Arizona", []int{2007, 2009, 2014}, []models.Player{}},
		{1611661319, "LVA", "Aces", 1997, "Las Vegas", "Las Vegas Aces", "Nevada", []int{2022, 2023}, []models.Player{}},
		{1611661320, "LAS", "Sparks", 1997, "Los Angeles", "Los Angeles Sparks", "California", []int{2001, 2002, 2016}, []models.Player{}},
		{1611661321, "DAL", "Wings", 1998, "Dallas", "Dallas Wings", "Texas", []int{2003, 2006, 2008}, []models.Player{}},
		{1611661322, "WAS", "Mystics", 1998, "Washington", "Washington Mystics", "District of Columbia", []int{2019}, []models.Player{}},
		{1611661323, "CON", "Sun", 1999, "Connecticut", "Connecticut Sun", "Connecticut", []int{}, []models.Player{}},
		{1611661324, "MIN", "Lynx", 1999, "Minnesota", "Minnesota Lynx", "Minnesota", []int{2011, 2013, 2015, 2017}, []models.Player{}},
		{1611661325, "IND", "Fever", 2000, "Indiana", "Indiana Fever", "Indiana", []int{2012}, []models.Player{}},
		{1611661328, "SEA", "Storm", 2000, "Seattle", "Seattle Storm", "Washington", []int{2004, 2010, 2018, 2020}, []models.Player{}},
		{1611661329, "CHI", "Sky", 2005, "Chicago", "Chicago Sky", "Illinois", []int{2021}, []models.Player{}},
		{1611661330, "ATL", "Dream", 2008, "Atlanta", "Atlanta Dream", "Georgia", []int{}, []models.Player{}},
	}
}

// GetWNBATeamsWithPlayers returns a hardcoded list of WNBA teams with their Roster
func GetWNBATeamsWithPlayers() Teams {
	wnbaTeams := GetWNBATeams()
	players := models.GetAllWNBAPlayers()
	if players == nil || len(players) == 0 {
		return wnbaTeams
	}
	for _, player := range players {
		fmt.Println(player.TeamID)
		wnbaTeams.GetTeamByID(player.TeamID).addRosterMember(player)
	}
	fmt.Println(wnbaTeams)
	return wnbaTeams
}

// GetNNBATeamsWithPlayers returns a hardcoded list of WNBA teams with their Roster
func GetNBATeamsWithPlayers() Teams {
	nbaTeams := GetNBATeams()
	players := models.GetAllNBAPlayers()

	if players == nil || len(players) == 0 {
		return nbaTeams
	}
	for _, player := range players {
		if v, ok := weirdNameMap[player.Name]; ok {
			player.Name = v
		} else {
			if strings.HasSuffix(player.Name, "Jr.") {
				// Remove just the trailing period
				player.Name = strings.TrimSuffix(player.Name, ".")
			}
		}
		team := nbaTeams.GetTeamByID(player.TeamID)
		if team != nil {
			team.addRosterMember(player)
		}

	}

	return nbaTeams
}

func GetNBAMatchups() []Matchup {
	gamesToday := models.GetNBAGamesToday()
	if gamesToday == nil {
		return nil
	}

	// Fetch the teams list once to prevent redundant calls
	nbaTeams := GetNBATeamsWithPlayers()
	var matchups []Matchup // Initialize slice to store matchups

	for _, m := range gamesToday {
		homeTeamID, err1 := m["HOME_TEAM_ID"].(json.Number).Int64()
		awayTeamID, err2 := m["VISITOR_TEAM_ID"].(json.Number).Int64()

		if err1 != nil || err2 != nil {
			fmt.Println("Error converting team IDs:", err1, err2)
			continue
		}

		homeTeam := nbaTeams.GetTeamByID(int(homeTeamID))
		awayTeam := nbaTeams.GetTeamByID(int(awayTeamID))

		if homeTeam == nil || awayTeam == nil {
			fmt.Printf("Warning: Could not find teams for match %d vs %d\n", homeTeamID, awayTeamID)
			continue
		}

		matchups = append(matchups, Matchup{
			HomeTeam: homeTeam,
			AwayTeam: awayTeam,
		})
	}

	return matchups
}

var weirdNameMap = map[string]string{
	"Nikola Jokić":       "Nikola Jokic",
	"Jimmy Butler III":   "Jimmy Butler",
	"Isaiah Stewart":     "Isaiah Stewart II",
	"Dennis Schröder":    "Dennis Schroder",
	"Kristaps Porziņģis": "Kristaps Porzingis",
	"Tim Hardaway Jr.":   "Tim Hardaway Jr",
	"Luka Dončić":        "Luka Doncic",
	"CJ McCollum":        "C.J. McCollum",
	"Danté Exum":         "Dante Exum",
	"Nick Smith Jr.":     "Nick Smith Jr",
	"RJ Barrett":         "R.J. Barrett",
	"Nic Claxton":        "Nicolas Claxton",
}

// GetNBAMatchupsWithOdds Get Today's Game and get acommpanying odds for the matchup
func GetNBAMatchupsWithOdds() []Matchup {
	gamesToday := models.GetNBAGamesToday()
	matchOdds, err := odds.GetPlayerProps()
	fullSeasonStats := models.GetAllNBAPlayerStatsFullSeason()
	if err != nil {
		return nil
	}
	if gamesToday == nil {
		return nil
	}

	// Fetch the teams list once to prevent redundant calls
	nbaTeams := GetNBATeamsWithPlayers()
	var matchups []Matchup // Initialize slice to store matchups

	for _, m := range gamesToday {
		fmt.Printf("%T", m["HOME_TEAM_ID"])
		homeTeamID := int(m["HOME_TEAM_ID"].(float64))
		awayTeamID := int(m["VISITOR_TEAM_ID"].(float64))
		homeTeam := nbaTeams.GetTeamByID(homeTeamID)
		awayTeam := nbaTeams.GetTeamByID(awayTeamID)

		if homeTeam == nil || awayTeam == nil {
			fmt.Printf("Warning: Could not find teams for match %d vs %d\n", homeTeamID, awayTeamID)
			continue
		}

		t := matchOdds.GetOddsByHomeAndAwayTeam(homeTeam.FullName, awayTeam.FullName)
		fmt.Println(t)
		for _, bookmaker := range t.Bookmakers {
			for _, market := range bookmaker.Markets {
				for _, outcome := range market.Outcomes {
					if homeTeam.GetPlayerByName(outcome.Description) != nil {
						player := homeTeam.GetPlayerByName(outcome.Description).SetOutcome(bookmaker.Key, market.Key, outcome.Name, outcome.Point, outcome.Price)
						logs := fullSeasonStats.GetPlayerGameLog(player.PlayerID)
						if logs != nil && len(logs) > 0 {
							player.SetCurrentSeasonLogs(logs).SetOpponentAbbreviation(awayTeam.Abbreviation)
						}

					} else if awayTeam.GetPlayerByName(outcome.Description) != nil {
						player := awayTeam.GetPlayerByName(outcome.Description).SetOutcome(bookmaker.Key, market.Key, outcome.Name, outcome.Point, outcome.Price)
						logs := fullSeasonStats.GetPlayerGameLog(player.PlayerID)
						if logs != nil && len(logs) > 0 {
							player.SetCurrentSeasonLogs(logs).SetOpponentAbbreviation(homeTeam.Abbreviation)
						}

					} else {
						fmt.Println("No Player Found For", outcome.Description)
					}
				}

			}
		}

		matchups = append(matchups, Matchup{
			HomeTeam: homeTeam,
			AwayTeam: awayTeam,
		})
	}

	return matchups
}

func GetActivePlayerForToday() []models.Player {
	matchups := GetNBAMatchupsWithOdds()
	var players []models.Player
	for _, matchup := range matchups {
		players = append(players, matchup.AwayTeam.Roster...)
		players = append(players, matchup.HomeTeam.Roster...)
	}
	return players
}
