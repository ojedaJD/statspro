package nba

import (
	client "sports_api/globals/nba"
	helpers "sports_api/helpers/nba"
	endpoints "sports_api/urls/nba"
)

// CumulativeStatsTeam calls the NBA API and retrieves cumulative team statistics based on the provided filters.
//
// Parameters:
// - gameIDs (string): Comma-separated list of 10-digit GameIDs.
// - leagueID (string): "00" (NBA), "10" (WNBA), "20" (G-League).
// - season (string): Format "YYYY-YY" (e.g., "2023-24").
// - seasonType (string): "Regular Season", "Pre Season", "Playoffs", or "All Star".
// - teamID (string): The ID of the team.
//
// Returns:
// - (*client.NBAResponse): The JSON response from the API, decoded into a struct.
// - (error): Returns an error if validation fails or if the API request encounters an issue.
//
// Example Usage:
//
//	stats, err := CumeStatsTeam("0021700807", "00", "2023-24", "Regular Season", "1610612739")
//	if err != nil {
//	    log.Fatal("Error fetching cumulative team stats:", err)
//	}
//	fmt.Println(stats)
func CumulativeStatsTeam(gameIDs, leagueID, season, seasonType, teamID string) (*client.NBAResponse, error) {
	// Validate input parameters
	if err := validateCumulativeStatsTeamParams(gameIDs, leagueID, season, seasonType, teamID); err != nil {
		return nil, err
	}

	params := map[string]string{
		"GameIDs":    gameIDs,
		"LeagueID":   leagueID,
		"Season":     season,
		"SeasonType": seasonType,
		"TeamID":     teamID,
	}

	return client.NBASession.NBAGetRequest(endpoints.CumulativeStatsTeam, params, "", nil)
}

// validateCumeStatsTeamParams ensures all input parameters are valid.
func validateCumulativeStatsTeamParams(gameIDs, leagueID, season, seasonType, teamID string) error {
	// Validate GameIDs using regex pattern
	if valid, err := helpers.ValidateGameIDs(gameIDs); !valid {
		return err
	}

	// Validate LeagueID
	if valid, err := helpers.ValidateLeagueID(leagueID); !valid {
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

	// Validate TeamID
	if valid, err := helpers.ValidateTeamID(teamID); !valid {
		return err
	}

	return nil
}
