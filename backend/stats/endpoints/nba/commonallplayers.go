package nba

import (
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
