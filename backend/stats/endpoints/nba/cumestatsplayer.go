package nba

import (
	client "sports_api/globals/nba"
	helpers "sports_api/helpers/nba"
	endpoints "sports_api/urls/nba"
)

// CumulativeStatsPlayer calls the NBA API to retrieve cumulative player statistics.
//
// Parameters:
// - gameIDs (string): One or more game IDs, separated by commas.
// - leagueID (string): The league identifier (e.g., "00" for NBA).
// - playerID (string): The player's unique identifier.
// - season (string): The season in "YYYY-YY" format (e.g., "2019-20").
// - seasonType (string): The type of season ("Regular Season", "Pre Season", "Playoffs", "All Star").
//
// Returns:
// - (*client.NBAResponse): The JSON response from the API, decoded into a struct.
// - (error): Returns an error if validation fails or if the API request encounters an issue.
//
// Example Usage:
//
//	stats, err := CumeStatsPlayer("0021700807", "00", "2544", "2019-20", "Regular Season")
//	if err != nil {
//	    log.Fatal("Error fetching cumulative player stats:", err)
//	}
//	fmt.Println(stats)
func CumulativeStatsPlayer(gameIDs, leagueID, playerID, season, seasonType string) (*client.NBAResponse, error) {
	if err := validateCumulativeStatsPlayerParams(gameIDs, leagueID, playerID, season, seasonType); err != nil {
		return nil, err
	}

	params := map[string]string{
		"GameIDs":    gameIDs,
		"LeagueID":   leagueID,
		"PlayerID":   playerID,
		"Season":     season,
		"SeasonType": seasonType,
	}

	return client.NBASession.NBAGetRequest(endpoints.CumulativeStatsPlayer, params, "", nil)
}

// validateCumulativeStatsPlayerParams ensures all input parameters are valid.
func validateCumulativeStatsPlayerParams(gameIDs, leagueID, playerID, season, seasonType string) error {
	// Validate GameIDs using regex pattern
	if valid, err := helpers.ValidateGameIDs(gameIDs); !valid {
		return err
	}

	// Validate LeagueID
	if valid, err := helpers.ValidateLeagueID(leagueID); !valid {
		return err
	}

	// Validate PlayerID
	if valid, err := helpers.ValidatePlayerID(playerID); !valid {
		return err
	}

	// Validate Season
	if valid, err := helpers.ValidateSeason(season); !valid {
		return err
	}

	// Validate SeasonType
	if valid, err := helpers.ValidateSeasonType(seasonType); !valid {
		return err
	}

	return nil
}
