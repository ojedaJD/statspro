package nba

import (
	"fmt"
	client "sports_api/globals/nba"
	helpers "sports_api/helpers/nba"
	endpoints "sports_api/urls/nba"
)

// AssistLeaders calls the NBA API and retrieves assist leader statistics based on the provided filters.
//
// Parameters:
// - leagueID (string):
//   - "00" → NBA (default)
//   - "10" → WNBA
//   - "20" → G-League
//
// - season (string): The season in "YYYY-YY" format (e.g., "2023-24").
// - seasonType (string): Specifies "Regular Season", "Playoffs", etc.
// - perMode (string): Specifies whether the data should be in "PerGame", "Totals", etc.
// - topX (int): Number of top players to retrieve (e.g., 10, 50, 100).
//
// Returns:
// - (*client.NBAResponse): The JSON response from the API, decoded into a struct.
// - (error): Returns an error if validation fails or if the API request encounters an issue.
//
// Example Usage:
//
//	leaders, err := AssistLeaders("00", "2023-24", "Regular Season", "PerGame", 10)
//	if err != nil {
//	    log.Fatal("Error fetching assist leaders:", err)
//	}
//	fmt.Println(leaders)
func AssistLeaders(leagueID, season, seasonType, perMode string, topX int) (*client.NBAResponse, error) {
	if err := validateAssistLeadersParams(leagueID, season, seasonType, perMode, topX); err != nil {
		return nil, err
	}

	params := map[string]string{
		"LeagueID":   leagueID,
		"Season":     season,
		"SeasonType": seasonType,
		"PerMode":    perMode,
		"TopX":       fmt.Sprintf("%d", topX),
	}

	return client.NBASession.NBAGetRequest(endpoints.AssistLeaders, params, "", nil)
}

// validateAssistLeadersParams ensures all input parameters are valid.
func validateAssistLeadersParams(leagueID, season, seasonType, perMode string, topX int) error {
	// Validate leagueID
	if valid, err := helpers.ValidateLeagueID(leagueID); !valid {
		return err
	}

	// Validate season format ("YYYY-YY")
	if valid, err := helpers.ValidateSeason(season); !valid {
		return err
	}

	// Validate seasonType
	if valid, err := helpers.ValidateSeasonType(seasonType); !valid {
		return err
	}

	// Validate perMode
	if valid, err := helpers.ValidatePerMode(perMode); !valid {
		return err
	}

	// Validate topX (must be a positive integer)
	if valid, err := helpers.IsPositive(topX); !valid {
		return err
	}

	return nil
}
