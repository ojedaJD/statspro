package nba

import (
	"encoding/json"
	"errors"
	"fmt"
	client "sports_api/globals/nba"
	helpers "sports_api/helpers/nba"
	endoints "sports_api/urls/nba"
)

// CommonAllPlayers calls the NBA API and retrieves all player data based on the provided filters.
//
// Parameters:
// - isOnlyCurrentSeason (int):
//   - `1` → Fetch only players from the current season.
//   - `0` → Fetch all players, including historical ones.
//
// - leagueID (string):
//   - `"00"` → NBA (default)
//   - `"10"` → WNBA
//   - `"20"` → G-League
//
// - season (string):
//   - Format: `"YYYY-YY"` (e.g., `"2023-24"` for the 2023-2024 season).
//
// Returns:
// - (interface{}): The JSON response from the API, decoded into an interface{}.
// - (error): Returns an error if validation fails or if the API request encounters an issue.
//
// Example Usage:
//
//	players, err := CommonAllPlayers(1, "00", "2023-24")
//	if err != nil {
//	    log.Fatal("Error fetching players:", err)
//	}
//	fmt.Println(players)
func CommonAllPlayers(isOnlyCurrentSeason int, leagueID, season string) (*client.NBAResponse, error) {

	if err := validateCommonAllPlayersParams(isOnlyCurrentSeason, leagueID, season); err != nil {
		return nil, err
	}

	params := map[string]string{
		"IsOnlyCurrentSeason": fmt.Sprintf("%d", isOnlyCurrentSeason),
		"LeagueID":            leagueID,
		"Season":              season,
	}

	return client.NBASession.NBAGetRequest(endoints.CommonAllPlayer, params, "", nil)
}

type Player struct {
	Name                 string                          `json:"DISPLAY_FIRST_LAST"`
	FromYear             string                          `json:"FROM_YEAR"`
	PlayerID             int                             `json:"PERSON_ID"`
	IsActive             int                             `json:"ROSTERSTATUS"`
	TeamAbbreviation     string                          `json:"TEAM_ABBREVIATION"`
	TeamCity             string                          `json:"TEAM_CITY"`
	TeamID               int                             `json:"TEAM_ID"`
	TeamName             string                          `json:"TEAM_NAME"`
	ToYear               string                          `json:"TO_YEAR"`
	Odds                 map[string]map[string][]Outcome `json:"odds,omitempty"`
	CurrentSeasonLogs    BaseGameLogSlice                // OutcomeType -> BookMaker -> Outcome
	OpponentAbbreviation string
}

func (r *Player) SetOpponentAbbreviation(str string) {
	r.OpponentAbbreviation = str
}

// SetOutcome method to add/update an outcome for a player
func (p *Player) SetOutcome(bookmaker, outcomeType, name string, point float64, price int) *Player {
	// Initialize Odds map if nil

	if p.Odds == nil {
		p.Odds = make(map[string]map[string][]Outcome)
	}

	// Initialize nested map for the bookmaker if nil
	if p.Odds[outcomeType] == nil {
		p.Odds[outcomeType] = make(map[string][]Outcome)
	}

	// Create a new outcome
	outcome := Outcome{
		Name:  name,
		Point: point,
		Price: price,
	}

	// Append the outcome to the existing slice
	p.Odds[outcomeType][bookmaker] = append(p.Odds[outcomeType][bookmaker], outcome)

	return p
}

func (p *Player) SetCurrentSeasonLogs(seasonLog BaseGameLogSlice) *Player {
	if len(seasonLog) == 0 || seasonLog == nil {
		fmt.Println("No logs found for ", p.Name)
		return nil
	}
	p.CurrentSeasonLogs = seasonLog
	return p

}

type Outcome struct {
	Name  string  `json:"name"`
	Point float64 `json:"point"`
	Price int     `json:"price"`
}

// Convert American odds to implied probability
func americanOddsToProbability(odds int) float64 {
	if odds > 0 {
		return float64(100) / float64(odds+100)
	}
	return float64(-odds) / float64(-odds+100)
}

// Calculate the optimal stake distribution for arbitrage betting
func calculateArbitrageBets(price1, price2 int, totalBankroll float64) (bool, float64, float64, float64) {
	prob1 := americanOddsToProbability(price1)
	prob2 := americanOddsToProbability(price2)
	fmt.Println(prob1, prob2)
	// Arbitrage condition: 1/prob1 + 1/prob2 < 1
	arbValue := (1/prob1 + 1/prob2)
	if arbValue >= 1 {
		return false, 0, 0, 0 // No arbitrage opportunity
	}

	// Optimal stake calculation
	stake1 := (totalBankroll * (1 / prob1)) / arbValue
	stake2 := (totalBankroll * (1 / prob2)) / arbValue
	profitPercentage := (1 - arbValue) * 100

	return true, stake1, stake2, profitPercentage
}

func GetAllNBAPlayers() []Player {
	players, err := CommonAllPlayers(1, "00", "2024-25")
	if err != nil {
		return nil
	}

	dict2, err := players.GetNormalizedDict2()

	if err != nil {
		return nil
	}
	currentSeasonPlayers := dict2["CommonAllPlayers"]
	jsonData, err := json.Marshal(currentSeasonPlayers)
	var player []Player
	err = json.Unmarshal(jsonData, &player)
	if err != nil {
		fmt.Println(err)
	}
	return player

}

func GetAllWNBAPlayers() []Player {
	players, err := CommonAllPlayers(1, "10", "2024-25")
	if err != nil {
		return nil
	}

	dict2, err := players.GetNormalizedDict2()

	if err != nil {
		return nil
	}
	currentSeasonPlayers := dict2["CommonAllPlayers"]
	jsonData, err := json.Marshal(currentSeasonPlayers)
	var player []Player
	err = json.Unmarshal(jsonData, &player)
	if err != nil {
		fmt.Println(err)
	}
	return player

}

// validateCommonAllPlayersParams ensures all input parameters are valid.
func validateCommonAllPlayersParams(isOnlyCurrentSeason int, leagueID, season string) error {
	// Validate isOnlyCurrentSeason (must be 0 or 1)
	if isOnlyCurrentSeason != 0 && isOnlyCurrentSeason != 1 {
		return errors.New("invalid value for IsOnlyCurrentSeason: must be 0 (all players) or 1 (current season only)")
	}

	// Validate leagueID using helper function
	if valid, err := helpers.ValidateLeagueID(leagueID); !valid {
		return err
	}

	// Validate season format using helper function
	if valid, err := helpers.ValidateSeason(season); !valid {
		return err
	}

	return nil
}
