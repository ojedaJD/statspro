package nba

import (
	"fmt"
	client "sports_api/globals/nba"
	helpers "sports_api/helpers/nba"
	endoints "sports_api/urls/nba"
)

// AllTimeLeadersGrids calls the NBA API and retrieves all-time leader statistics based on the provided filters.
//
// Parameters:
// - leagueID (string):
//   - "00" → NBA (default)
//   - "10" → WNBA
//   - "20" → G-League
//
// - perMode (string): Specifies whether the data should be in "Totals", "PerGame", etc.
// - seasonType (string): Specifies "Regular Season", "Playoffs", etc.
// - topX (int): Number of top players to retrieve (e.g., 10, 50, 100).
//
// Returns:
// - (*client.NBAResponse): The JSON response from the API, decoded into a struct.
// - (error): Returns an error if validation fails or if the API request encounters an issue.
//
// Example Usage:
//
//	leaders, err := AllTimeLeadersGrids("00", "Totals", "Regular Season", 10)
//	if err != nil {
//	    log.Fatal("Error fetching all-time leaders:", err)
//	}
//	fmt.Println(leaders)
func AllTimeLeadersGrids(leagueID, perMode, seasonType string, topX int) (*client.NBAResponse, error) {
	if err := validateAllTimeLeadersParams(leagueID, perMode, seasonType, topX); err != nil {
		return nil, err
	}

	params := map[string]string{
		"LeagueID":   leagueID,
		"PerMode":    perMode,
		"SeasonType": seasonType,
		"TopX":       fmt.Sprintf("%d", topX),
	}

	return client.NBASession.NBAGetRequest(endoints.AllTimeLeadersGrids, params, "", nil)
}

// validateAllTimeLeadersParams ensures all input parameters are valid.
func validateAllTimeLeadersParams(leagueID, perMode, seasonType string, topX int) error {
	if valid, err := helpers.ValidateLeagueID(leagueID); !valid {
		return err
	}

	if valid, err := helpers.ValidatePerMode(perMode); !valid {
		return err
	}

	if valid, err := helpers.ValidateSeasonType(seasonType); !valid {
		return err
	}

	if valid, err := helpers.IsPositive(topX); !valid {
		return err
	}

	return nil
}
